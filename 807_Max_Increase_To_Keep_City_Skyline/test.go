package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMaxs := make([]int, len(grid))
	colMaxs := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			e := grid[i][j]
			if e > rowMaxs[i] {
				rowMaxs[i] = e
			}
			if e > colMaxs[j] {
				colMaxs[j] = e
			}
		}
	}
	increments := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			e := grid[i][j]
			increments += (min(rowMaxs[i], colMaxs[j]) - e)
		}
	}
	return increments
}

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}
