package main

import (
	"app/main/backend_examples/mvc"
	"app/main/data_structures/bstree"
	"fmt"
)

type MyInt int

func (a MyInt) Less(b interface{}) bool {
	return a < b.(MyInt)
}

func main() {
	tree := new(bstree.BSTree)
	tree.Insert(MyInt(13))
	tree.Insert(MyInt(6))
	tree.Insert(MyInt(8))
	tree.Insert(MyInt(15))
	tree.Insert(MyInt(7))
	fmt.Printf("%T -> %v\n", tree, tree.InorderList())

	mvc.CreateAPI()
}
