package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	for level := 0; level < len(matrix)/2; level++ {
		topBoundary := level
		leftBoundary := level
		bottomBoundary := len(matrix) - 1 - level
		rightBoundary := len(matrix) - 1 - level
		for times := 0; times < bottomBoundary-topBoundary; times++ {
			// left
			mostLeftTop := matrix[topBoundary][leftBoundary]
			for row := topBoundary; row < bottomBoundary; row++ {
				matrix[row][leftBoundary] = matrix[row+1][leftBoundary]
			}

			// bottom
			for col := leftBoundary; col < rightBoundary; col++ {
				matrix[bottomBoundary][col] = matrix[bottomBoundary][col+1]
			}

			// right
			for row := bottomBoundary; row > topBoundary; row-- {
				matrix[row][rightBoundary] = matrix[row-1][rightBoundary]
			}

			// top
			for col := rightBoundary; col > leftBoundary+1; col-- {
				matrix[topBoundary][col] = matrix[topBoundary][col-1]
			}
			matrix[topBoundary][leftBoundary+1] = mostLeftTop
		}
	}
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
