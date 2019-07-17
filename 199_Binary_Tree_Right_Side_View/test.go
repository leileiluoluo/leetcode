package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrderTraversal(root *TreeNode, depth int, r *[]int) {
	if len(*r) < depth {
		*r = append(*r, root.Val)
	}
	if nil != root.Right {
		postOrderTraversal(root.Right, depth+1, r)
	}
	if nil != root.Left {
		postOrderTraversal(root.Left, depth+1, r)
	}
}

func rightSideView(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}
	r := []int{root.Val}
	postOrderTraversal(root, 1, &r)
	return r
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	fmt.Println(rightSideView(root))
}
