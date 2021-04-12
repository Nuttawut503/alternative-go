package huffman_coding

import (
	"container/heap"
)

type Node struct {
	item                byte
	freq                int
	leftNode, rightNode *Node
}

type NodeList []Node

func (list NodeList) Len() int {
	return len(list)
}

func (list NodeList) Less(i, j int) bool {
	return list[i].freq < list[j].freq
}

func (list NodeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list *NodeList) Push(item interface{}) {
	*list = append(*list, item.(Node))
}

func (list *NodeList) Pop() interface{} {
	old := *list
	item := old[len(old)-1]
	old[len(old)-1] = Node{}
	*list = old[:len(old)-1]
	return item
}

func BuildHuffman(text []byte, freq []int) map[byte]string {
	sli := make(NodeList, len(text))

	for i := 0; i < len(text); i++ {
		sli[i] = Node{
			item: text[i],
			freq: freq[i],
		}
	}

	heap.Init(&sli)

	for len(sli) > 1 {
		min1 := heap.Pop(&sli).(Node)
		min2 := heap.Pop(&sli).(Node)
		newNode := Node{
			item:      '-',
			freq:      min1.freq + min2.freq,
			leftNode:  &min1,
			rightNode: &min2,
		}
		heap.Push(&sli, newNode)
	}

	codes := map[byte]string{}
	root := heap.Pop(&sli).(Node)
	readCode(&root, codes, "")
	return codes
}

func readCode(node *Node, codes map[byte]string, code string) {
	if node == nil {
		return
	}

	if node.item != '-' {
		codes[node.item] = code
	}

	readCode(node.leftNode, codes, code+"0")
	readCode(node.rightNode, codes, code+"1")
}
