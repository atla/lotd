package ldb

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

// Helper ... asd
type Helper struct {
	Path string
	db   *leveldb.DB
}

// NewHelper ... creates new login manager
func NewHelper(path string) *Helper {

	db, err := leveldb.OpenFile(path, nil)

	if err != nil {
		log.Panic("Could not open database " + path)
	}

	return &Helper{
		Path: path,
		db:   db,
	}
}

// CLose ... close the db
func (helper *Helper) Close() {
	defer helper.db.Close()
}

// Put ... puts a value
func (helper *Helper) Put(key string, value []byte) error {
	return helper.db.Put([]byte(key), value, nil)
}

// Get ... gets the value
func (helper *Helper) Get(key string) ([]byte, error) {
	data, err := helper.db.Get([]byte(key), nil)
	return data, err
}

// Delete ... deletes a value
func (helper *Helper) Delete(key string) error {
	return helper.db.Delete([]byte(key), nil)
}
