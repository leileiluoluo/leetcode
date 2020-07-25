package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if nil == head || nil == head.Next ||
		m == n {
		return head
	}

	step := n - m
	var leftTail *ListNode
	p := head
	for m > 1 {
		leftTail = p
		p = p.Next
		m--
	}

	q := p.Next
	p.Next = nil
	midTail := p
	for step > 0 {
		r := q.Next
		q.Next = p
		p = q
		q = r
		step--
	}

	if nil == leftTail {
		midTail.Next = q
		return p
	}

	leftTail.Next = p
	midTail.Next = q
	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	reverseBetween(head, 1, 4)
}
