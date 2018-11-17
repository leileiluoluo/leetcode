package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type CircularListNode struct {
	*ListNode
	Pre *CircularListNode
}

func reorderList(head *ListNode) {
	// validation
	if nil == head {
		return
	}

	// build circular list
	tail := &CircularListNode{head, nil}
	for node := head.Next; nil != node; node = node.Next {
		currNode := &CircularListNode{node, tail}
		tail = currNode
	}

	// re order
	for p, q := head, tail; ; {
		if p == q.ListNode {
			p.Next = nil
			break
		}
		if p.Next == q.ListNode {
			p.Next.Next = nil
			break
		}
		x, y := p, q
		p, q = p.Next, q.Pre
		x.Next = y.ListNode
		y.Next = p
	}
}

func main() {
	head := &ListNode{Val: 1}
	for node, v := head, 2; v <= 5; v++ {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	reorderList(head)
	for n := head; nil != n; n = n.Next {
		fmt.Println(n)
	}
}
