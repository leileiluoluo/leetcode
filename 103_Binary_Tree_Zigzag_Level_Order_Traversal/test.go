package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if nil == root {
		return [][]int{}
	}

	var allVals [][]int
	nodes := []*TreeNode{root}
	level := 0
	for len(nodes) > 0 {
		var vals []int
		tmp := nodes[:]
		nodes = []*TreeNode{}
		for _, p := range tmp {
			if 0 == level%2 {
				vals = append(vals, p.Val)
			} else {
				vals = append([]int{p.Val}, vals...)
			}
			if nil != p.Left {
				nodes = append(nodes, p.Left)
			}
			if nil != p.Right {
				nodes = append(nodes, p.Right)
			}
		}
		allVals = append(allVals, vals)
		level++
	}
	return allVals
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}
