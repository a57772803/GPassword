package util

import "github.com/syndtr/goleveldb/leveldb"

func GetDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile("test.db", nil)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *leveldb.DB) error {
	err := db.Close()
	return err
}
