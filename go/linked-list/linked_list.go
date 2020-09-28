// Go exercise linked-list
package linkedlist

import "errors"

// Node represents a node in a linked list.
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// Return the next node in the list.
func (node *Node) Next() *Node {
	return node.next
}

// Return the previous node in the list.
func (node *Node) Prev() *Node {
	return node.prev
}

// A doubly linked list.
type List struct {
	head *Node
	tail *Node
}

// Error for operations that are not allowed on an empty list.
var ErrEmptyList = errors.New("Empty list")

// Create a new list and populate it with the given values.
func NewList(values ...interface{}) *List {
	list := List{}
	for _, value := range values {
		list.PushBack(value)
	}
	return &list
}

// Add a value to the front of the list.
func (list *List) PushFront(value interface{}) {
	node := Node{Val: value, next: list.head, prev: nil}
	if list.head != nil {
		list.head.prev = &node
	}
	list.head = &node
	if list.tail == nil {
		list.tail = &node
	}
}

// Add a value to the back of the list.
func (list *List) PushBack(value interface{}) {
	node := Node{Val: value, next: nil, prev: list.tail}
	if list.tail != nil {
		list.tail.next = &node
	}
	list.tail = &node
	if list.head == nil {
		list.head = &node
	}
}

// Remove the first value from the list and return it.
func (list *List) PopFront() (interface{}, error) {
	if list.head == nil {
		return nil, ErrEmptyList
	}
	value := list.head.Val
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	return value, nil
}

// Remove the last value from the list and return it.
func (list *List) PopBack() (interface{}, error) {
	if list.tail == nil {
		return nil, ErrEmptyList
	}
	value := list.tail.Val
	list.tail = list.tail.prev
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = nil
	}
	return value, nil
}

// Reverse the list in-place.
func (list *List) Reverse() *List {
	if list.head == nil {
		return list
	}
	front := list.head
	back := list.tail
	for {
		if front == back {
			return list
		}
		front.Val, back.Val = back.Val, front.Val
		front = front.next
		if front == back {
			return list
		}
		back = back.prev
	}
}

// Return the first node in the list.
func (list *List) First() *Node {
	return list.head
}

// Return the last node in the list.
func (list *List) Last() *Node {
	return list.tail
}
