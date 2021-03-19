package main

import (
	"app/main/bstree"
	dq "app/main/deque"
	"fmt"
)

func main() {
	tree := new(bstree.BSTree)
	tree.Insert(13)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(7)
	fmt.Printf("%T -> %v\n", tree, tree.InorderList())

	deque := new(dq.Deque)
	deque.PushFront(10)
	deque.PushFront(2)
	deque.PopBack()
	deque.PushFront(3)
	deque.PushBack(4)
	fmt.Printf("%T -> %v length: %d\n", deque, deque.ToList(), deque.Len())
}
