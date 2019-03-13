package storage

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger"
	"github.com/mrwinstead/knv/observation"
	"github.com/pkg/errors"
)

const (
	IndexNamePrimaryKey = "PrimaryKey"

	sinkNameOperationDelete = "BadgerBackedStorageDelete"
	sinkNameOperationGet    = "BadgerBackedStorageGet"
	sinkNameOperationPut    = "BadgerBackedStoragePut"
	sinkNameBadgerLog       = "BackerBackedStorageLibraryTextLog"

	operationNameDelete = "Delete"
	operationNameGet    = "Get"
	operationNamePut    = "Put"
)

var (
	_                    MultiKeyValue = &badgerBackedStore{}
	badgerDefaultOptions               = badger.DefaultOptions
)

type BadgerStorageObservable struct {
	Operation string

	Table string
	Index string

	Key string

	Value          []byte
	AdditionalKeys map[string]string

	Error string
}

type badgerKey struct {
	Table, Index, Key []byte
}

type badgerBackedStore struct {
	rootDirectoryName string

	// tableCodecs is a codec per Table
	tableCodecs map[string]Codec

	sinkDelete observation.Sink
	sinkGet    observation.Sink
	sinkPut    observation.Sink

	rootTable *badger.DB
}

// NewBadgerBackedStore create a new MultiKeyValue store using Badger as the
// underlying storage engine
func NewBadgerBackedStore(rootDirectoryName string, tableCodecs map[string]Codec,
	manager observation.Manager) (MultiKeyValue, error) {
	opts := badgerDefaultOptions
	opts.Dir = rootDirectoryName
	opts.ValueDir = rootDirectoryName
	rootTable, openErr := badger.Open(opts)
	if nil != openErr {
		return nil, openErr
	}

	created := &badgerBackedStore{
		rootTable:   rootTable,
		tableCodecs: tableCodecs,
		sinkGet:     manager.Sink(sinkNameOperationGet),
		sinkDelete:  manager.Sink(sinkNameOperationDelete),
		sinkPut:     manager.Sink(sinkNameOperationPut),
	}
	return created, nil
}

func (b *badgerBackedStore) Put(table, key string,
	additionalKeys map[string]string, value interface{}) error {
	obs := &BadgerStorageObservable{
		Operation:      operationNamePut,
		Key:            key,
		Table:          table,
		AdditionalKeys: additionalKeys,
	}
	defer b.sinkPut.Submit(nil, obs)

	primaryKeyFormatted := b.formatKey(table, IndexNamePrimaryKey,
		key)
	formattedAdditionalKeys := make([][]byte, 0, len(additionalKeys))
	for index, key := range additionalKeys {
		formattedAdditionalKeys = append(formattedAdditionalKeys,
			b.formatKey(table, index, key))
	}

	serializedValueBuf := bytes.Buffer{}
	serializeBufErr := gob.NewEncoder(&serializedValueBuf).Encode(value)
	if nil != serializeBufErr {
		return serializeBufErr
	}
	obs.Value = serializedValueBuf.Bytes()

	txnErr := b.rootTable.Update(func(txn *badger.Txn) error {
		setErr := txn.Set(primaryKeyFormatted, serializedValueBuf.Bytes())
		if nil != setErr {
			return setErr
		}

		for _, key := range formattedAdditionalKeys {
			setErr = txn.Set(key, primaryKeyFormatted)
			if nil != setErr {
				return setErr
			}
		}

		return nil
	})
	if nil != txnErr {
		obs.Error = txnErr.Error()
		return txnErr
	}

	return nil
}

func (b *badgerBackedStore) Get(table, index, key string) (interface{},
	error) {
	var value interface{}
	obs := &BadgerStorageObservable{
		Operation: operationNameGet,
		Key:       key,
		Table:     table,
		Index:     index,
	}
	defer b.sinkPut.Submit(nil, obs)

	formattedKey := b.formatKey(table, index, key)
	viewErr := b.rootTable.View(func(txn *badger.Txn) error {
		item, metadataGetErr := txn.Get(formattedKey)
		if badger.ErrKeyNotFound == metadataGetErr ||
			item.IsDeletedOrExpired() {
			err := &ErrNotFound{
				Table: table,
				Index: index,
				Key:   key,
			}
			return err
		} else if nil != metadataGetErr {
			err := errors.Wrap(metadataGetErr,
				"could not fetch item metadata from database")
			return err
		}

		valueSerialized, valueCopyErr := item.ValueCopy(nil)
		if nil != valueCopyErr {
			err := errors.Wrap(valueCopyErr,
				"could not fetch item data from database")
			return err
		}
		obs.Value = valueSerialized

		valueUnmarshalled, unmarshalErr := b.tableCodecs[table].
			Unmarshal(valueSerialized)
		if nil != unmarshalErr {
			err := errors.Wrapf(unmarshalErr,
				"could not unmarshal value from database on "+
					"Table %s, Index %s, Key %v",
				table, index, key)
			return err
		}

		value = valueUnmarshalled
		return nil
	})

	if IsErrNotFound(viewErr) {
		obs.Error = viewErr.Error()
		return nil, viewErr

	} else if nil != viewErr {
		err := errors.Wrap(viewErr, "could not complete read")
		obs.Error = err.Error()
		return nil, err
	}

	return value, nil
}

func (b *badgerBackedStore) Delete(table, index, key string) error {
	obs := &BadgerStorageObservable{
		Operation: operationNameDelete,
		Key:       key,
		Table:     table,
		Index:     index,
	}
	defer b.sinkPut.Submit(nil, obs)

	formattedKey := b.formatKey(table, index, key)
	deleteErr := b.rootTable.Update(func(txn *badger.Txn) error {
		item, metadataGetErr := txn.Get(formattedKey)
		if badger.ErrKeyNotFound == metadataGetErr ||
			item.IsDeletedOrExpired() {
			err := &ErrNotFound{
				Table: table,
				Index: index,
				Key:   key,
			}
			return err
		} else if nil != metadataGetErr {
			err := errors.Wrap(metadataGetErr,
				"could not fetch item metadata from database")
			return err
		}

		metadataDeleteErr := txn.Delete(formattedKey)
		if nil != metadataDeleteErr {
			err := errors.Wrap(metadataDeleteErr,
				"could not fetch item metadata from database")
			return err
		}
		return nil
	})

	if IsErrNotFound(deleteErr) {
		obs.Error = deleteErr.Error()
		return deleteErr
	}

	if nil != deleteErr {
		obs.Error = deleteErr.Error()
		return deleteErr
	}

	return deleteErr
}

func (b *badgerBackedStore) formatKey(tableName, indexName, key string) []byte {
	k := badgerKey{
		Table: []byte(tableName),
		Index: []byte(indexName),
		Key:   []byte(key),
	}
	combined := bytes.Buffer{}
	encodeErr := gob.NewEncoder(&combined).Encode(k)
	if nil != encodeErr {
		err := errors.Wrap(encodeErr, "error encoding Key for "+
			"database, this should not happen")
		panic(err)
	}
	return combined.Bytes()
}
