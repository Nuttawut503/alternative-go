package bstree

import (
	"reflect"
	"testing"
)

func TestBSTree(t *testing.T) {
	t.Run("Check if bstree works properly", func(t *testing.T) {
		tree := new(BSTree)
		tree.Insert(3)
		tree.Insert(1)
		tree.Insert(0)
		tree.Insert(2)
		tree.Insert(4)
		tree.Insert(5)
		got := tree.InorderList()
		expect := []int{0, 1, 2, 3, 4, 5}
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})
}
