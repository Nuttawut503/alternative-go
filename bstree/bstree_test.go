package bstree

import (
	"reflect"
	"testing"
)

type IntVar int

func (a IntVar) Less(b interface{}) bool {
	return a < b.(IntVar)
}

func TestBSTree(t *testing.T) {
	t.Run("Check if bstree works properly", func(t *testing.T) {
		tree := new(BSTree)
		tree.Insert(IntVar(3))
		tree.Insert(IntVar(1))
		tree.Insert(IntVar(0))
		tree.Insert(IntVar(2))
		tree.Insert(IntVar(4))
		tree.Insert(IntVar(5))
		treeList := tree.InorderList()
		got := make([]int, len(treeList))
		for i, v := range treeList {
			got[i] = int(v.(IntVar))
		}
		expect := []int{0, 1, 2, 3, 4, 5}
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %v but got %v", expect, got)
		}
	})
}
