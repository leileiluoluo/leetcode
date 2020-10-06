package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if 0 == len(preorder) {
		return nil
	}

	root := preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if root == inorder[i] {
			break
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  buildTree(preorder[1:i+1], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder[i+1:]),
	}
}

func (tree *TreeNode) String() string {
	if nil == tree {
		return ""
	}
	return fmt.Sprintf("[left: %s, val: %d, right: %s]", tree.Left, tree.Val, tree.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	tree := buildTree(preorder, inorder)
	fmt.Println(tree)
}
