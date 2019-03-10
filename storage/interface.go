package storage

type InstanceFactory func() interface{}

// MultiKeyValue represents the data storage layer
type MultiKeyValue interface {
	Put(tableName, primaryKey string, additionalKeys map[string]string,
		value interface{}) error
	Get(tableName, key, indexName string) (interface{}, error)
	Delete(tableName, key, indexname string) error
}

type Codec interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte) (interface{}, error)
}
