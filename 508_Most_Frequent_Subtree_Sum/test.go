package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeSum(root *TreeNode, frequents map[int]int) int {
	if nil == root {
		return 0
	}
	sum := root.Val + treeSum(root.Left, frequents) + treeSum(root.Right, frequents)
	if v, ok := frequents[sum]; ok {
		frequents[sum] = v + 1
	} else {
		frequents[sum] = 1
	}
	return sum
}

func findFrequentTreeSum(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	frequents := make(map[int]int)
	treeSum(root, frequents)

	vMax := 1
	maxMap := make(map[int][]int)
	for k, v := range frequents {
		if v < vMax {
			continue
		}
		if v > vMax {
			vMax = v
		}
		maxMap[vMax] = append(maxMap[vMax], k)
	}
	return maxMap[vMax]
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: -5,
		},
	}
	fmt.Println(findFrequentTreeSum(root))
}
