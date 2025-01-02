package kv

import (
	"project/btree"
)

type KV struct {
	Path string // file name // internals
	fd   int
	tree btree.BTree
	// more ...
}

func (db *KV) Open() error {

}
func (db *KV) Get(key []byte) ([]byte, bool) {
	return db.tree.Get(key)
}
func (db *KV) Set(key []byte, val []byte) error {
	db.tree.Insert(key, val)
	return updateFile(db)
}
func (db *KV) Del(key []byte) (bool, error) {
	deleted := db.tree.Delete(key)
	return deleted, updateFile(db)
}
