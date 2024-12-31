package test

import (
	"project/btree"
	"testing"
)

func TestBtreeRead(t *testing.T) {
	c := btree.NewC()
	c.Add("1", "1")
	c.Add("2", "2")

	val, ok := c.Read("1")

	if !ok {
		t.Error("Read fail")
	}

	if val != c.Ref["1"] {
		t.Errorf("Read fail: expected %v, got %v", c.Ref["1"], val)
	}
}

func TestReadMissingKey(t *testing.T) {
	c := btree.NewC()

	val, ok := c.Read("fdsa")
	if ok {
		t.Errorf("read missing key and got value %s", val)
	}
}
