package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func spiralMatrixIII(r int, c int, r0 int, c0 int) [][]int {
	walk := make([][]int, r*c)
	i, j := r0, c0
	index := 0

	walk[index] = []int{i, j}
	index++
	circle := 1
	for r0+circle < r || r0-circle >= 0 ||
		c0+circle < c || c0-circle >= 0 {
		// down direction
		i = max(0, r0-circle+1)
		j = c0 + circle
		for j < c && i < r0+circle && i < r {
			walk[index] = []int{i, j}
			index++
			i++
		}

		// left direction
		i = r0 + circle
		j = min(c-1, c0+circle)
		for i < r && j > c0-circle && j >= 0 {
			walk[index] = []int{i, j}
			index++
			j--
		}

		// up direction
		i = min(r-1, r0+circle)
		j = c0 - circle
		for j >= 0 && i > r0-circle && i >= 0 {
			walk[index] = []int{i, j}
			index++
			i--
		}

		// right direction
		i = r0 - circle
		j = max(0, c0-circle)
		for i >= 0 && j <= c0+circle && j < c {
			walk[index] = []int{i, j}
			index++
			j++
		}
		circle++
	}
	return walk
}

func main() {
	r, c := 5, 6
	r0, c0 := 1, 4
	fmt.Println(spiralMatrixIII(r, c, r0, c0))
}
