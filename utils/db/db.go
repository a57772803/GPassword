package db

import "github.com/syndtr/goleveldb/leveldb"

func GetDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *leveldb.DB) error {
	err := db.Close()
	return err
}
