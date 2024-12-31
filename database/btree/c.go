package btree

import (
	"project/utils"
	"unsafe"
)

type C struct {
	tree  BTree
	Ref   map[string]string
	pages map[uint64]BNode
}

func NewC() *C {
	pages := map[uint64]BNode{}
	return &C{
		tree: BTree{
			get: func(ptr uint64) []byte {
				node, ok := pages[ptr]
				utils.Assert(ok, "Can't read allocated data")
				return node
			},
			new: func(node []byte) uint64 {
				utils.Assert(BNode(node).nbytes() <= BTREE_PAGE_SIZE, "new node exceed max size")
				ptr := uint64(uintptr(unsafe.Pointer(&node[0])))
				utils.Assert(pages[ptr] == nil, "pointer already been assigned")
				pages[ptr] = node
				return ptr
			},
			del: func(ptr uint64) {
				utils.Assert(pages[ptr] != nil, "try to de-allocate a pointer that is not occupied")
				delete(pages, ptr)
			},
		},
		Ref:   map[string]string{},
		pages: pages,
	}
}

func (c *C) Read(key string) (string, bool) {
	val, ok := c.tree.Read([]byte(key))
	return string(val), ok
}

func (c *C) Add(key string, val string) {
	c.tree.Insert([]byte(key), []byte(val))
	c.Ref[key] = val
}

func (c *C) Del(key string) {
	c.tree.Delete([]byte(key))
	delete(c.Ref, key)
}
