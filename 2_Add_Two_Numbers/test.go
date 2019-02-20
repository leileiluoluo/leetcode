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

func (l *ListNode) String() string {
	var vals []string
	for nil != l {
		vals = append(vals, strconv.Itoa(l.Val))
		l = l.Next
	}
	return strings.Join(vals, " -> ")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	p, q, r := l1, l2, l
	carry := 0
	for nil != p || nil != q || carry > 0 {
		v := carry
		if nil != p {
			v += p.Val
			p = p.Next
		}
		if nil != q {
			v += q.Val
			q = q.Next
		}
		carry, r.Val = v/10, v%10
		if nil != p || nil != q || carry > 0 {
			r.Next = &ListNode{}
			r = r.Next
		}
	}
	return l
}

func main() {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}
	l := addTwoNumbers(l1, l2)
	fmt.Println(l)
}
