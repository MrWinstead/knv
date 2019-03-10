package storage

import (
	"context"
	"github.com/mrwinstead/knv/observation"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func buildTestBadgerBackedStore(t *testing.T, tableCodecs map[string]Codec) (
	MultiKeyValue, observation.Manager, func()) {
	tmpDir, _ := ioutil.TempDir(os.TempDir(), "badger-")
	cleanupFunc := func() {
		if err := os.RemoveAll(tmpDir); nil != err {
			panicErr := errors.Wrapf(err,
				"could not clean up badger database directory at %s",
				tmpDir)
			panic(panicErr)
		}
	}

	obsManager := observation.NewChannelManager(100)

	mkv, buildErr := NewBadgerBackedStore(tmpDir, tableCodecs, obsManager)
	assert.NoError(t, buildErr)
	assert.NotNil(t, mkv)
	return mkv, obsManager, cleanupFunc
}

func TestNewBadgerBackedStore(t *testing.T) {
	mkv, _, cleanupFunc := buildTestBadgerBackedStore(t, nil)
	defer cleanupFunc()
	assert.NotNil(t, mkv)
}

func TestBadgerBackedStore_Put(t *testing.T) {
	tableName := gofakeit.UUID()
	tableCodecs := map[string]Codec{
		tableName: NewGobCodec(func() interface{} {
			return &gofakeit.CreditCardInfo{}
		}),
	}

	mkv, _, cleanupFunc := buildTestBadgerBackedStore(t, tableCodecs)
	defer cleanupFunc()
	primaryKey := gofakeit.UUID()
	creditCard := gofakeit.CreditCard()

	putErr := mkv.Put(tableName, primaryKey, nil, creditCard)
	assert.NoError(t, putErr)

	fetched, getErr := mkv.Get(tableName, IndexNamePrimaryKey, primaryKey)
	assert.NoError(t, getErr)

	assert.Equal(t, creditCard, fetched)
}

func TestBadgerBackedStore_Put_Observation(t *testing.T) {
	tableName := gofakeit.UUID()
	tableCodecs := map[string]Codec{
		tableName: NewGobCodec(func() interface{} {
			return &gofakeit.CreditCardInfo{}
		}),
	}

	mkv, obsManager, cleanupFunc := buildTestBadgerBackedStore(t, tableCodecs)
	defer cleanupFunc()

	recorder := observation.NewRecordingObserver()
	obsManager.Register(recorder)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		runErr := obsManager.Run(ctx)
		assert.NoError(t, runErr)
	}()

	primaryKey := gofakeit.UUID()
	creditCard := gofakeit.CreditCard()

	putErr := mkv.Put(tableName, primaryKey, nil, creditCard)
	assert.NoError(t, putErr)

	_, getErr := mkv.Get(tableName, IndexNamePrimaryKey, primaryKey)
	assert.NoError(t, getErr)

	events := recorder.Events()
	assert.NotEmpty(t, events)
	assert.Equal(t, 1, len(events),
		"expected a single recorded event")
	evt := events[0].Value.(*BadgerStorageObservable)
	assert.Equal(t, "Put", evt.Operation)
}

func TestBadgerBackedStore_Get_NotFound(t *testing.T) {
	tableName := gofakeit.UUID()
	tableCodecs := map[string]Codec{
		tableName: NewGobCodec(func() interface{} {
			return &gofakeit.CreditCardInfo{}
		}),
	}

	mkv, _, cleanupFunc := buildTestBadgerBackedStore(t, tableCodecs)
	defer cleanupFunc()
	primaryKey := gofakeit.UUID()

	fetched, getErr := mkv.Get(tableName, IndexNamePrimaryKey, primaryKey)
	assert.Nil(t, fetched)
	assert.Error(t, getErr)
	assert.IsType(t, &ErrNotFound{}, getErr,
		"expected ErrNorFound as returned error")
}

func TestBadgerBackedStore_Delete_NotFound(t *testing.T) {
	tableName := gofakeit.UUID()
	tableCodecs := map[string]Codec{
		tableName: NewGobCodec(func() interface{} {
			return &gofakeit.CreditCardInfo{}
		}),
	}

	mkv, _, cleanupFunc := buildTestBadgerBackedStore(t, tableCodecs)
	defer cleanupFunc()
	primaryKey := gofakeit.UUID()

	deleteErr := mkv.Delete(tableName, IndexNamePrimaryKey, primaryKey)
	assert.Error(t, deleteErr)
	assert.IsType(t, &ErrNotFound{}, deleteErr,
		"expected ErrNorFound as returned error")
}
