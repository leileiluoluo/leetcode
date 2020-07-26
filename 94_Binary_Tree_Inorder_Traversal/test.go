package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	vals := []int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]

		if nil != node.Left {
			nodes = append(nodes, node.Left)
			node.Left = nil
			continue
		}

		vals = append(vals, node.Val)
		nodes = nodes[:len(nodes)-1]
		if nil != node.Right {
			nodes = append(nodes, node.Right)
		}
	}

	return vals
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		},
	}

	fmt.Println(inorderTraversal(root))
}
