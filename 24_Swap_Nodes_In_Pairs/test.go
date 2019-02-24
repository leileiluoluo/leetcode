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
	for nil != head {
		nodes = append(nodes, strconv.Itoa(head.Val))
		head = head.Next
	}
	return strings.Join(nodes, ", ")
}

func swapPairs(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	p := head
	head = head.Next
	q := p.Next
	for nil != q {
		r := q.Next
		q.Next = p
		if nil == r {
			p.Next = nil
			break
		}
		if nil == r.Next {
			p.Next = r
			r.Next = nil
			break
		}
		p.Next = r.Next
		p = r
		q = p.Next
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{Val: 4}}}}
	head = swapPairs(head)
	fmt.Println(head)
}
