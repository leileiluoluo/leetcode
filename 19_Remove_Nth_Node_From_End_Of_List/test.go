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

func (node *ListNode) String() string {
	var nodes []string
	for nil != node {
		nodes = append(nodes, strconv.Itoa(node.Val))
		node = node.Next
	}
	return strings.Join(nodes, ", ")
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	for ; n > 0; n-- {
		p = p.Next
	}
	if nil == p {
		return head.Next
	}
	for nil != p.Next {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return head
}

func main() {
	head := &ListNode{Val: 1}
	fmt.Println(removeNthFromEnd(head, 1))
}
