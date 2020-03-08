package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	largestVals := []int{}
	children := []*TreeNode{root}
	for len(children) > 0 {
		tmp := children[:]
		children = []*TreeNode{}
		largest := -(1 << 32)
		for _, child := range tmp {
			if child.Val > largest {
				largest = child.Val
			}
			if nil != child.Left {
				children = append(children, child.Left)
			}
			if nil != child.Right {
				children = append(children, child.Right)
			}
		}
		largestVals = append(largestVals, largest)
	}
	return largestVals
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(largestValues(root))
}
