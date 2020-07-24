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
	q := reverseList(head.Next)
	p.Next = nil
	r := q
	for nil != r.Next {
		r = r.Next
	}
	r.Next = p
	return q
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
