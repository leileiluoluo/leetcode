package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	// special
	if 0 == len(matrix) {
		return []int{}
	}

	// standard
	var elements []int

	leftBoundary, rightBundary := 0, len(matrix[0])-1
	topBoundary, bottomBoundary := 0, len(matrix)-1

	for leftBoundary < rightBundary &&
		topBoundary < bottomBoundary {
		// top
		for j := leftBoundary; j < rightBundary; j++ {
			elements = append(elements, matrix[topBoundary][j])
		}

		// right
		for i := topBoundary; i < bottomBoundary; i++ {
			elements = append(elements, matrix[i][rightBundary])
		}

		// bottom
		for j := rightBundary; j > leftBoundary; j-- {
			elements = append(elements, matrix[bottomBoundary][j])
		}

		// left
		for i := bottomBoundary; i > topBoundary; i-- {
			elements = append(elements, matrix[i][leftBoundary])
		}

		leftBoundary++
		rightBundary--
		topBoundary++
		bottomBoundary--
	}

	if leftBoundary == rightBundary &&
		topBoundary <= bottomBoundary {
		for i := topBoundary; i <= bottomBoundary; i++ {
			elements = append(elements, matrix[i][leftBoundary])
		}
		return elements
	}

	if topBoundary == bottomBoundary &&
		leftBoundary <= rightBundary {
		for j := leftBoundary; j <= rightBundary; j++ {
			elements = append(elements, matrix[topBoundary][j])
		}
	}

	return elements
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))
}
