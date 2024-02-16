package ozone

import (
	"errors"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// OzoneWrapper interface defines the set of methods to interact with a LevelDB instance.
type OzoneWrapper interface {
	Put(key any, value any) error
	Get(key any) ([]byte, error)
	Delete(key any) error
	NewIterator(slice *util.Range, ro *opt.ReadOptions) iterator.Iterator
	WriteBatch(batch *leveldb.Batch) error
	Close() error
}

// Entry interface should be implemented by types that need to be stored in LevelDB.
// It requires methods for serialization and deserialization.
type Entry interface {
	ToBytes() []byte
	FromBytes(data []byte)
}

// Ozone is a wrapper around LevelDB to provide simplified access to its functions.
type Ozone struct {
	db *leveldb.DB
}

// Put inserts or updates a value in the database with the specified key.
func (wrapper *Ozone) Put(key any, value any) error {
	keyBytes, err := toBytes(key)
	if err != nil {
		return err
	}
	valueBytes, err := toBytes(value)
	if err != nil {
		return err
	}

	log.Printf("Putting key: %s, value: %v\n", string(keyBytes), valueBytes)
	if err = wrapper.db.Put(keyBytes, valueBytes, nil); err != nil {
		return err
	}

	return nil
}

// Get retrieves a value from the database based on the specified key.
func (wrapper *Ozone) Get(key any) ([]byte, error) {
	keyBytes, err := toBytes(key)
	if err != nil {
		return nil, err
	}

	data, err := wrapper.db.Get(keyBytes, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete removes a value from the database associated with the specified key.
func (wrapper *Ozone) Delete(key any) error {
	keyBytes, err := toBytes(key)
	if err != nil {
		return err
	}

	if err = wrapper.db.Delete(keyBytes, nil); err != nil {
		return err
	}

	return nil
}

// NewIterator creates an iterator for traversing a subset or the entire database.
func (wrapper *Ozone) NewIterator(slice *util.Range, ro *opt.ReadOptions) iterator.Iterator {
	return wrapper.db.NewIterator(slice, ro)
}

// WriteBatch applies multiple operations (put/delete) in a batch.
func (wrapper *Ozone) WriteBatch(batch *leveldb.Batch) error {
	return wrapper.db.Write(batch, nil)
}

// toBytes converts a given data to a byte slice. It supports types implementing Entry and []byte.
func toBytes(data any) ([]byte, error) {
	switch v := data.(type) {
	case Entry:
		return v.ToBytes(), nil
	case []byte:
		return v, nil
	default:
		return nil, errors.New("data.(type) NOT Entry or []Byte")
	}
}

// NewOzone creates and returns a new Ozone instance for interacting with LevelDB at the given location.
func NewOzone(location string) (*Ozone, error) {
	ldb, err := leveldb.OpenFile(location, nil)
	if err != nil {
		return nil, err
	}
	return &Ozone{db: ldb}, nil
}

// Close safely closes the LevelDB database connection.
func (wrapper *Ozone) Close() error {
	if err := wrapper.db.Close(); err != nil {
		return err
	}
	return nil
}
