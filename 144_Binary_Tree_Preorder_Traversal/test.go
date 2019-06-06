package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var vals []int
	if nil == root {
		return vals
	}

	// represent for right nodes waiting for traversal
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		for p := node; nil != p; p = p.Left {
			vals = append(vals, p.Val)
			if node == p {
				nodes = nodes[:len(nodes)-1]
			}
			if nil != p.Right {
				nodes = append(nodes, p.Right)
			}
		}
	}
	return vals
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(preorderTraversal(root))
}
