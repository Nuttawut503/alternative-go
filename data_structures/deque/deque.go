package deque

type Deque struct {
	head   *dqNode
	tail   *dqNode
	length int
}

type dqType interface{}

type dqNode struct {
	item     dqType
	prevNode *dqNode
	nextNode *dqNode
}

func (ll *Deque) PushFront(item dqType) {
	newNode := new(dqNode)
	newNode.item = item
	if ll.head == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		newNode.nextNode = ll.head
		ll.head.prevNode = newNode
		ll.head = newNode
	}
	ll.length++
}

func (ll *Deque) PushBack(item dqType) {
	newNode := new(dqNode)
	newNode.item = item
	if ll.tail == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		newNode.prevNode = ll.tail
		ll.tail.nextNode = newNode
		ll.tail = newNode
	}
	ll.length++
}

func (ll *Deque) PopFront() {
	if ll.length < 2 {
		ll.Clear()
	} else {
		ll.head = ll.head.nextNode
		ll.head.prevNode = nil
		ll.length--
	}
}

func (ll *Deque) PopBack() {
	if ll.length < 2 {
		ll.Clear()
	} else {
		ll.tail = ll.tail.prevNode
		ll.tail.nextNode = nil
		ll.length--
	}
}

func (ll *Deque) Len() int {
	return ll.length
}

func (ll *Deque) ToList() []interface{} {
	list := make([]interface{}, 0)
	ll.head.saveToList(&list)
	return list
}

func (node *dqNode) saveToList(list *[]interface{}) {
	if node != nil {
		*list = append(*list, node.item)
		node.nextNode.saveToList(list)
	}
}

func (ll *Deque) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}
