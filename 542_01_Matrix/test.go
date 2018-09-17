package main

import (
	"fmt"
	"math/rand"
)

const MaxInt = 999999999

func steps(x1, y1, x2, y2 int) int {
	abs := func(a, b int) int {
		c := a - b
		if c < 0 {
			return -c
		}
		return c
	}
	return abs(x1, x2) + abs(y1, y2)
}

func minSteps(m [][]int, x, y int) int {
	minSteps := 100000
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if 0 == m[i][j] {
				steps := steps(x, y, i, j)
				if steps < minSteps {
					minSteps = steps
				}
			}
		}
	}
	return minSteps
}

func findMinSteps(m [][]int, x, y int, radius int) int {
	minSteps := MaxInt
	for i := x - radius; i >= 0 && i <= x+radius && x < len(m); i++ {
		for j := y - radius; j >= 0 && j <= y+radius && y < len(m[i]); j++ {
			if 0 == m[i][j] {
				minSteps = steps(x, y, i, j)
			}
		}
	}
	return minSteps
}

func containsZero(sumGraph [][]int, x, y int, radius int) bool {
	xMin, yMin, xMax, yMax := x-radius, y-radius, x+radius, y+radius
	if xMin < 0 {
		xMin = 0
	}
	if yMin < 0 {
		yMin = 0
	}
	if xMax > len(sumGraph)-1 {
		xMax = len(sumGraph) - 1
	}
	if yMax > len(sumGraph)-1 {
		yMax = len(sumGraph) - 1
	}

	sumAllNotZeros := (xMax - xMin) * (yMax - yMin + 1)
	sum := sumAllNotZeros
	if 0 == xMin && 0 == yMin {
		sum = sumGraph[xMax][yMax]
	} else if 0 == xMin && yMin > 0 {
		sum = sumGraph[xMax][yMax] - sumGraph[xMax - 1][yMin - 1]
	} else if xMin > 0 && 0 == yMin {
		sum = sumGraph[xMax][yMax] - sumGraph[xMin - 1][yMax - 1]
	} else {
		sum = sumGraph[xMax][yMax] - sumGraph[xMin][yMax] - sumGraph[xMax][yMin] + sumGraph[xMin][yMin]
	}
	return sum < sumAllNotZeros
}

func findMinStepsAround(m, sumGraph [][]int, x, y int, len int) int {
	return 1
}

func sumGraph(m [][]int) [][]int {
	sumGraph := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		sumGraph[i] = make([]int, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			sumGraph[i][j] = 0
			if 0 == j && 0 == i {
				sumGraph[i][j] = m[i][j]
			} else if 0 == j && i > 0 {
				sumGraph[i][j] = sumGraph[i-1][j] + m[i][j]
			} else if j > 0 && 0 == i {
				sumGraph[i][j] = sumGraph[i][j-1] + m[i][j]
			} else {
				sumGraph[i][j] = sumGraph[i][j-1] + sumGraph[i-1][j] - sumGraph[i-1][j-1] + m[i][j]
			}
		}
	}
	return sumGraph
}

func updateMatrix(m [][]int) [][]int {
	sumGraph := sumGraph(m)
	contains := containsZero(sumGraph, 5, 5, 2)
	fmt.Println(contains)
	return sumGraph
}

func print(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if len(m[i])-1 == j {
				fmt.Println(m[i][j])
			} else {
				fmt.Printf("%d ", m[i][j])
			}
		}
	}
}

func main() {
	var m [][]int
	for i := 0; i < 10; i++ {
		var r []int
		for j := 0; j < 10; j++ {
			r = append(r, rand.Intn(2))
		}
		m = append(m, r)
	}
	print(m)
	m = updateMatrix(m)
	fmt.Println("-----------------------")
	print(m)
}
