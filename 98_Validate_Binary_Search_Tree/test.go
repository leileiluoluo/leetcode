package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode, preVal *int) bool {
	if nil == root {
		return true
	}
	ok := preOrderTraversal(root.Left, preVal)
	if !ok {
		return false
	}
	if root.Val <= *preVal {
		return false
	}
	*preVal = root.Val
	return preOrderTraversal(root.Right, preVal)
}

func isValidBST(root *TreeNode) bool {
	preVal := -(1 << 32)
	return preOrderTraversal(root, &preVal)
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isValidBST(root))
}
