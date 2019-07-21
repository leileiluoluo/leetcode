package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type withNo struct {
	*TreeNode
	No int
}

func isCompleteTree(root *TreeNode) bool {
	var nodes []withNo
	preNo := 0
	no := 1
	nodes = append(nodes, withNo{root, no})
	for len(nodes) > 0 {
		node := nodes[0]
		if preNo+1 != node.No {
			return false
		}
		nodes = nodes[1:]
		no++
		if nil != node.Left {
			nodes = append(nodes, withNo{node.Left, no})
		}
		no++
		if nil != node.Right {
			nodes = append(nodes, withNo{node.Right, no})
		}
		preNo = node.No
	}
	return true
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
				Val: 7,
			},
		},
	}
	fmt.Println(isCompleteTree(root))
}
