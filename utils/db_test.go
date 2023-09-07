package utils

import (
	"fmt"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestGetDB(t *testing.T) {
	db, err := leveldb.OpenFile("testdb", nil)

	fmt.Println(db)
	fmt.Print(err)
}

func TestDbInsert(t *testing.T) {
	db, err := leveldb.OpenFile("./tizi365.db", nil)
	db.Put([]byte("test"), []byte("123"), nil)
	db.Put([]byte("test"), []byte("444"), nil)
	data, err := db.Get([]byte("test"), nil)
	err = db.Delete([]byte("test"), nil)
	fmt.Println(data)
	fmt.Print(err)
}
