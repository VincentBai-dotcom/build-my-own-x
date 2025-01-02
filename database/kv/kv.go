package kv

import (
	"fmt"
	"os"
	"path"
	"project/btree"
	"syscall"
)

type KV struct {
	Path string // file name
	// internals
	fd   int
	tree btree.BTree
	// more ...
}

func (db *KV) Open() error {
	db.tree.Get = db.pageRead   // read a page
	db.tree.New = db.pageAppend // apppend a page
	db.tree.Del = func(uint64) {}
}

func (db *KV) Get(key []byte) ([]byte, bool) {
	return db.tree.Read(key)
}
func (db *KV) Set(key []byte, val []byte) error {
	db.tree.Insert(key, val)
	return updateFile(db)
}
func (db *KV) Del(key []byte) (bool, error) {
	deleted := db.tree.Delete(key)
	return deleted, updateFile(db)
}

func updateFile(db *KV) error {
	// 1. Write new nodes.
	if err := writePages(db); err != nil {
		return err
	}
	// 2. `fsync` to enforce the order between 1 and 3.
	if err := syscall.Fsync(db.fd); err != nil {
		return err
	}
	// 3. Update the root pointer atomically.
	if err := updateRoot(db); err != nil {
		return err
	}
	// 4. `fsync` to make everything persistent.
	return syscall.Fsync(db.fd)
}

func createFileSync(file string) (int, error) {
	// obtain the directory fd
	flags := os.O_RDONLY | syscall.O_DIRECTORY
	dirfd, err := syscall.Open(path.Dir(file), flags, 0o644)
	if err != nil {
		return -1, fmt.Errorf("open directory: %w", err)
	}
	defer syscall.Close(dirfd)
	// open or create the file
	flags = os.O_RDWR | os.O_CREATE
	fd, err := syscall.Openat(dirfd, path.Base(file), flags, 0o644)
	if err != nil {
		return -1, fmt.Errorf("open file: %w", err)
	}
	// fsync the directory
	if err = syscall.Fsync(dirfd); err != nil {
		_ = syscall.Close(fd) // may leave an empty file return -1, fmt.Errorf("fsync directory: %w", err)
	}
	return fd, nil
}
