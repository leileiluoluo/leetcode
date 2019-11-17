package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if 0 == len(preorder) {
		return nil
	}

	val := preorder[0]
	preorder = preorder[1:]
	i := 0
	for i < len(preorder) && preorder[i] < val {
		i++
	}
	return &TreeNode{
		val,
		bstFromPreorder(preorder[:i]),
		bstFromPreorder(preorder[i:]),
	}
}

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}
	fmt.Println(bstFromPreorder(preorder))
}
