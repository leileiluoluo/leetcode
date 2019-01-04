package main

import "fmt"

func setZeroes(matrix [][]int) {
	zeroCols := make(map[int]int)
	for i := range matrix {
		rowHasZero := false
		for j := range matrix[i] {
			if 0 == matrix[i][j] {
				rowHasZero = true
				zeroCols[j] = 1
			}
			if _, ok := zeroCols[j]; ok {
				for k := 0; k <= i; k++ {
					matrix[k][j] = 0
				}
			}
		}
		if rowHasZero {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
