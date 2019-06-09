package main

import "fmt"

type node struct {
	val  int
	pre  *node
	next *node
}

type MyCircularDeque struct {
	front *node
	rear  *node
	len   int
	cap   int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{&node{}, &node{}, 0, k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.cap == this.len {
		return false
	}
	if 0 == this.len {
		this.front = &node{val: value}
		this.rear = this.front
	} else {
		front := this.front
		this.front = &node{val: value, next: front}
		front.pre = this.front
	}
	this.len++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.cap == this.len {
		return false
	}
	if 0 == this.len {
		this.rear = &node{val: value}
		this.front = this.rear
	} else {
		rear := this.rear
		this.rear = &node{val: value, pre: rear}
		rear.next = this.rear
	}
	this.len++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if 0 == this.len {
		return false
	}
	next := this.front.next
	if nil == next {
		this.front = nil
		this.rear = nil
	} else {
		this.front = this.front.next
		this.front.pre = nil
	}
	this.len--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if 0 == this.len {
		return false
	}
	pre := this.rear.pre
	if nil == pre {
		this.rear = nil
		this.front = nil
	} else {
		this.rear = this.rear.pre
		this.rear.next = nil
	}
	this.len--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if 0 == this.len {
		return -1
	}
	return this.front.val
}

func (this *MyCircularDeque) GetRear() int {
	if 0 == this.len {
		return -1
	}
	return this.rear.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return 0 == this.len
}

func (this *MyCircularDeque) IsFull() bool {
	return this.cap == this.len
}

func main() {
	deque := Constructor(3)
	fmt.Println(deque.InsertLast(1))
	fmt.Println(deque.InsertLast(2))
	fmt.Println(deque.InsertFront(3))
	fmt.Println(deque.InsertFront(4))
	fmt.Println(deque.GetRear())
	fmt.Println(deque.IsFull())
	fmt.Println(deque.DeleteLast())
	fmt.Println(deque.InsertFront(4))
	fmt.Println(deque.GetFront())
}
