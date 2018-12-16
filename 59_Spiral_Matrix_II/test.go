package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for row := 0; row < len(matrix); row++ {
		matrix[row] = make([]int, n)
	}

	v := 1
	topBoundary, rightBoundary, bottomBoundary,
		leftBoundary := 0, len(matrix[0]), len(matrix), 0
	for topBoundary < bottomBoundary {
		// top
		for col := leftBoundary; col < rightBoundary; col++ {
			matrix[topBoundary][col] = v
			v++
		}
		topBoundary++

		// right
		for row := topBoundary; row < bottomBoundary; row++ {
			matrix[row][rightBoundary-1] = v
			v++
		}
		rightBoundary--

		// bottom
		for col := rightBoundary - 1; col >= leftBoundary; col-- {
			matrix[bottomBoundary-1][col] = v
			v++
		}
		bottomBoundary--

		// left
		for row := bottomBoundary - 1; row >= topBoundary; row-- {
			matrix[row][leftBoundary] = v
			v++
		}
		leftBoundary++
	}

	return matrix
}

func main() {
	matrix := generateMatrix(3)
	fmt.Println(matrix)
}
