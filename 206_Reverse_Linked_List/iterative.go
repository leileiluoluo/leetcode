package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	p := head
	q := p.Next
	p.Next = nil
	for nil != q {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}
	return p
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

	reverseList(head)
}
