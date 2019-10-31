package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	nodes := []*TreeNode{root}
	maxDepth := -1
	val := 0
	for depth := 0; len(nodes) > 0; depth++ {
		copy := nodes[:]
		nodes = []*TreeNode{}
		for _, node := range copy {
			if depth > maxDepth {
				maxDepth = depth
				val = node.Val
			}
			if nil != node.Left {
				nodes = append(nodes, node.Left)
			}
			if nil != node.Right {
				nodes = append(nodes, node.Right)
			}
		}
	}
	return val
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(findBottomLeftValue(root))
}
