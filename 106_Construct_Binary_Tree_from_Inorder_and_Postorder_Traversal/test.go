package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if 0 == len(inorder) {
		return nil
	}

	val := postorder[len(postorder)-1]
	root := &TreeNode{Val: val}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == val {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}

func main() {
	inorder := []int{1, 2, 3}
	postorder := []int{1, 2, 3}
	buildTree(inorder, postorder)
}
