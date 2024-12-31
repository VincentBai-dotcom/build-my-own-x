package test

import (
	"project/btree"
	"testing"
)

func TestReadMissingValue(t *testing.T) {
	bTreeController := btree.NewC()

	_, ok := bTreeController.Read("fdsa")
	if ok {
		print("ok")
	}
}
