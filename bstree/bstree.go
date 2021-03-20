package bstree

type BSTree struct {
	root *bstreeNode
}

type bstreeType interface {
	Less(interface{}) bool
}

type bstreeNode struct {
	item      bstreeType
	leftNode  *bstreeNode
	rightNode *bstreeNode
}

func (tree *BSTree) Insert(item bstreeType) {
	if tree.root == nil {
		tree.root = new(bstreeNode)
		tree.root.item = item
	} else {
		tree.root.insert(item)
	}
}

func (node *bstreeNode) insert(item bstreeType) {
	if item.Less(node.item) {
		if node.leftNode == nil {
			node.leftNode = new(bstreeNode)
			node.leftNode.item = item
		} else {
			node.leftNode.insert(item)
		}
	} else if node.rightNode == nil {
		node.rightNode = new(bstreeNode)
		node.rightNode.item = item
	} else {
		node.rightNode.insert(item)
	}
}

func (tree *BSTree) InorderList() []interface{} {
	list := make([]interface{}, 0)
	tree.root.inorder(&list)
	return list
}

func (node *bstreeNode) inorder(list *[]interface{}) {
	if node != nil {
		node.leftNode.inorder(list)
		*list = append(*list, node.item)
		node.rightNode.inorder(list)
	}
}
