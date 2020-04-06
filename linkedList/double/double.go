package double

import "fmt"

type Ele struct {
	Key   string
	Value interface{}
}

type DoubleNode struct {
	Value Ele
	Pre   *DoubleNode
	Next  *DoubleNode
}

type Double struct {
	Header *DoubleNode
	Tail   *DoubleNode
	Length int
}

// time complexity 1
func (list *Double) Append(value Ele) {
	newNode := &DoubleNode{value, nil, nil}
	list.Length++

	if list.Header == nil {
		list.Header = newNode
		list.Tail = newNode
		return
	}

	list.Tail.Next = newNode
	newNode.Pre = list.Tail
}

func (list *Double) RemoveNode(node *DoubleNode) {

	if list.Header == node {
		list.Header = list.Header.Next
	}

	if list.Tail == node {
		list.Tail = list.Tail.Pre
		list.Tail.Next = nil
	}

	if list.Header.Next == nil {
		list.Tail = list.Header
	}

}

func (list *Double) UNShiftNode(node *DoubleNode) {

	if list.Header == nil {
		list.Header = node
		list.Tail = node
		return
	}

	node.Pre = nil
	node.Next = list.Header
	list.Header.Pre = node
	list.Header = node

}

func (t *Double) Traverse() {

	n := t.Header

	for {
		if n == nil {
			break
		}
		fmt.Println(n.Value)
		n = n.Next
	}

}
