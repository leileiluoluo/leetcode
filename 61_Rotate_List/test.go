package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	var nodes []string
	for p := head; nil != p; p = p.Next {
		nodes = append(nodes, strconv.Itoa(p.Val))
	}
	return strings.Join(nodes, " -> ")
}

func rotateRight(head *ListNode, k int) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	p, q := head, head
	len := 0

	// firstly, p move right k steps
Loop:
	for k > 0 {
		k--
		len++
		if nil == p.Next {
			// k is equals to len, do not move, return immediatly
			if 0 == k {
				return head
			}
			// k is larger than len, k mod len, then go back to the beginning
			k %= len
			p = head
			len = 0
			goto Loop
		}
		p = p.Next
	}

	// then, p/q move right together util p arriving at tail
	for nil != p.Next {
		p = p.Next
		q = q.Next
	}

	// re-build relations
	p.Next = head
	head = q.Next
	q.Next = nil

	return head
}

func main() {
	head := &ListNode{Val: 1}
	p := head
	for i := 2; i < 6; i++ {
		node := &ListNode{Val: i}
		p.Next = node
		p = p.Next
	}

	fmt.Println(head)
	head = rotateRight(head, 6)
	fmt.Println(head)
}
