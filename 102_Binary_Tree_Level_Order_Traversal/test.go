package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return [][]int{}
	}

	var vals [][]int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		currLevel := []int{}
		copy := nodes[:]
		nodes = []*TreeNode{}
		for _, node := range copy {
			currLevel = append(currLevel, node.Val)
			if nil != node.Left {
				nodes = append(nodes, node.Left)
			}
			if nil != node.Right {
				nodes = append(nodes, node.Right)
			}
		}
		vals = append(vals, currLevel)
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
		},
	}
	fmt.Println(levelOrder(root))
}
