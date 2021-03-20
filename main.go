package main

import (
	"app/main/bstree"
	dq "app/main/deque"
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

	deque := new(dq.Deque)
	deque.PushFront(10)
	deque.PushFront(2)
	deque.PopBack()
	deque.PushFront(3)
	deque.PushBack(4)
	fmt.Printf("%T -> %v length: %d\n", deque, deque.ToList(), deque.Len())

}
