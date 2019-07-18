package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generate(begin, end int) []*TreeNode {
	if begin > end {
		return []*TreeNode{nil}
	}
	if begin == end {
		return []*TreeNode{{Val: begin}}
	}
	var trees []*TreeNode
	for i := begin; i <= end; i++ {
		left := generate(begin, i-1)
		right := generate(i+1, end)
		for _, j := range left {
			for _, k := range right {
				root := &TreeNode{Val: i, Left: j, Right: k}
				trees = append(trees, root)
			}
		}
	}
	return trees
}

func generateTrees(n int) []*TreeNode {
	if 0 == n {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func main() {
	generateTrees(3)
}
