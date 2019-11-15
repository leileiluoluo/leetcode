package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	var vals []int
	leftNodes := []*TreeNode{root}
	for len(leftNodes) > 0 {
		node := leftNodes[len(leftNodes)-1]
		leftNodes = leftNodes[:len(leftNodes)-1]
		for nil != node {
			vals = append([]int{node.Val}, vals...)
			if nil != node.Left {
				leftNodes = append(leftNodes, node.Left)
			}
			node = node.Right
		}
	}
	return vals
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(postorderTraversal(root))
}
