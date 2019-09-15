package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	if nil != node {
		return "(Val: " + strconv.Itoa(node.Val) +
			", Next: " + node.Next.String() + ")"
	}
	return "nil"
}

func partition(head *ListNode, x int) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	var left, right, p, q *ListNode
	for ; nil != head; head = head.Next {
		if head.Val < x {
			if nil == left {
				left = &ListNode{Val: head.Val}
				p = left
			} else {
				p.Next = &ListNode{Val: head.Val}
				p = p.Next
			}
		} else {
			if nil == right {
				right = &ListNode{Val: head.Val}
				q = right
			} else {
				q.Next = &ListNode{Val: head.Val}
				q = q.Next
			}
		}
	}

	if nil == left {
		return right
	}
	if nil != right {
		p.Next = right
	}
	return left
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	fmt.Println(partition(head, 3))
}
