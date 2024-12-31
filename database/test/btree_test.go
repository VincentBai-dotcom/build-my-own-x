package test

import (
	"project/btree"
	"testing"
)

var bTreeController = btree.NewC()

func TestSomething(t *testing.T) {
	bTreeController.Add("1", "fdas")
}
