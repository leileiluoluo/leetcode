package main

import (
	"fmt"
	"strings"
)

// MyLinkedList is a struct of singly linked list
type MyLinkedList struct {
	head *Node // the head of the linked list
	tail *Node // the tail of the linked list
	len  int   // length of the linked list
}

// Node is used to construct linked list
type Node struct {
	val  int   // value of current node
	next *Node // point to next node
}

// Constructor is used to construct a empty list
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get is used to get the vaule of index-th node
// If index is invalid, -1 will be returned
// Otherwise, It returns the vaule of index-th node in the linked list
func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= list.len {
		return -1
	}

	p := list.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p.val
}

// AddAtHead is used to add a node of val at the head of the linked list
func (list *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, nil}

	if 0 == list.len {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head = node
	}

	list.len++
}

// AddAtTail is used to add a node of val at the tail of the linked list
func (list *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, nil}

	if 0 == list.len {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}

	list.len++
}

// AddAtIndex is used to add a node of value val bofore the index-th node of the linked list
// If index equals the length of the linked list, the node be added will be the tail of the list
// If index is invalid, do nothing
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > list.len {
		return
	}

	node := &Node{val, nil}
	if 0 == index { // add at head
		node.next = list.head
		list.head = node
		if 0 == list.len {
			list.tail = node
		}
	} else if list.len == index { // add at tail
		list.tail.next = node
		list.tail = node
	} else {
		p := list.head
		for i := 0; i < index-1; i++ {
			p = p.next
		}
		node.next = p.next
		p.next = node
	}

	list.len++
}

// DeleteAtIndex is used to delete the index-th node of the linked list
// If index is invalid, do nothing
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.len {
		return
	}

	if 0 == index { // delete head
		list.head = list.head.next
		if 1 == list.len {
			list.tail = nil
		}
	} else {
		p := list.head
		for i := 0; i < index-1; i++ {
			p = p.next
		}
		p.next = p.next.next
		if list.len-1 == index {
			list.tail = p
		}
	}

	list.len--
}

// String used to return the string of the linked list
func (list MyLinkedList) String() string {
	vals := []string{}
	p := list.head
	for nil != p {
		vals = append(vals, fmt.Sprintf("%d", p.val))
		p = p.next
	}
	return strings.Join(vals, ", ")
}

func main() {
	list := Constructor()

	list.AddAtIndex(0, 10)
	fmt.Println(list)

	list.AddAtIndex(0, 20)
	fmt.Println(list)

	list.AddAtIndex(1, 30)
	fmt.Println(list)

	fmt.Println(list.Get(0))
}
