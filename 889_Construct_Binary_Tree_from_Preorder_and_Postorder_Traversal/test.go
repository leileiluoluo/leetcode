package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if 0 == len(pre) {
		return nil
	}
	if 1 == len(pre) {
		return &TreeNode{Val: pre[0]}
	}

	root := &TreeNode{Val: pre[0]}
	pre = pre[1:]
	post = post[:len(post)-1]
	i := 0
	for i < len(post) {
		if pre[0] == post[i] {
			break
		}
		i++
	}
	root.Left = constructFromPrePost(pre[:i+1], post[:i+1])
	root.Right = constructFromPrePost(pre[i+1:], post[i+1:])
	return root
}

func main() {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	fmt.Println(constructFromPrePost(pre, post))
}
