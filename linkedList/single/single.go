package single

import (
	"fmt"
	"reflect"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type Single struct {
	Header *Node
	Tail   *Node
	Length int
}

func (list *Single) getNode(index int) *Node {
	if index < 0 {
		return nil
	}

	n := list.Header
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n

}

// time Complexity n
func (list *Single) GetValueAtIndex(index int) interface{} {
	n := list.getNode(index)
	return n.Value
}

func (list *Single) DeletValue(value interface{}) bool {

	currentNode := list.Header
	if reflect.DeepEqual(currentNode.Value, value) {
		list.Header = list.Header.Next
		if list.Header == nil {
			list.Tail = nil
		}

		list.Length--
		return true
	}
	for {
		if currentNode.Next == nil {
			break
		}
		if reflect.DeepEqual(currentNode.Next.Value, value) {
			currentNode.Next = currentNode.Next.Next
			if list.Header.Next == nil {
				list.Tail = list.Header
			}

			list.Length--
			return true
		}

		currentNode = currentNode.Next
	}

	return false

}

func (list *Single) DeleteNode(index int) bool {
	if index > list.Length-1 || list.Header == nil {
		return false

	}
	if index == 0 {
		list.Header = list.Header.Next
		if list.Header == nil {
			list.Tail = nil
		}
	} else {
		preNode := list.getNode(index - 1)
		preNode.Next = preNode.Next.Next
	}

	list.Length--
	if list.Length == 1 {
		list.Tail = list.Header
	}

	return true

}

// time complexity 1
func (list *Single) Append(value interface{}) {
	newNode := &Node{value, nil}
	list.Length++

	if list.Header == nil {
		list.Header = newNode
		list.Tail = newNode

		return
	}

	list.Tail.Next = newNode
	list.Tail = newNode
}

func (list *Single) Traverse() {
	n := list.Header
	for {
		if n == nil {
			break
		}
		fmt.Println(n, n.Value)
		n = n.Next
	}
}

//
func (list *Single) Reverse() {
	if list.Header == nil || list.Header.Next == nil {
		return
	}

	pre := list.Header
	current := list.Header.Next
	pre.Next = nil

	for {
		if current == nil {
			break
		}
		temp := current.Next
		current.Next = pre

		pre = current
		current = temp
	}

	list.Tail = list.Header
	list.Header = pre
}
