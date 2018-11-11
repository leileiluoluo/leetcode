package main

import "fmt"

// findOBoundaries return coordinates that value is 'O'
// and mark the coordinates visited
func findOBoundaries(board [][]byte, oVisited *[][]int) [][]int {
	var boundaries [][]int
	if 1 == len(board) {
		if 'O' == board[0][0] {
			p := []int{0, 0}
			boundaries = append(boundaries, p)
			*oVisited = append(*oVisited, p)
		}
		return boundaries
	}

	// left boundary
	for i := 0; i < len(board)-1; i++ {
		if 'O' == board[i][0] {
			p := []int{i, 0}
			boundaries = append(boundaries, p)
			*oVisited = append(*oVisited, p)
		}
	}

	// bottom boundary
	for j := 0; j < len(board[0])-1; j++ {
		if 'O' == board[len(board)-1][j] {
			p := []int{len(board) - 1, j}
			boundaries = append(boundaries, p)
			*oVisited = append(*oVisited, p)
		}
	}

	// right boundary
	for i := len(board) - 1; i > 0; i-- {
		if 'O' == board[i][len(board[0])-1] {
			p := []int{i, len(board[0]) - 1}
			boundaries = append(boundaries, p)
			*oVisited = append(*oVisited, p)
		}
	}

	// top boundary
	for j := len(board[0]) - 1; j > 0; j-- {
		if 'O' == board[0][j] {
			p := []int{0, j}
			boundaries = append(boundaries, p)
			*oVisited = append(*oVisited, p)
		}
	}

	return boundaries
}

func isVisited(i, j int, oVisited *[][]int) bool {
	for _, v := range *oVisited {
		if v[0] == i && v[1] == j {
			return true
		}
	}
	return false
}

func findOsAround(board [][]byte, i, j int, oVisited *[][]int) [][]int {
	var osAround [][]int

	// left
	if j > 0 {
		left := board[i][j-1]
		if !isVisited(i, j-1, oVisited) && 'O' == left {
			p := []int{i, j - 1}
			osAround = append(osAround, p)
			*oVisited = append(*oVisited, p)
		}
	}

	// down
	if i < len(board)-1 {
		down := board[i+1][j]
		if !isVisited(i+1, j, oVisited) && 'O' == down {
			p := []int{i + 1, j}
			osAround = append(osAround, p)
			*oVisited = append(*oVisited, p)
		}
	}

	// right
	if j < len(board[0])-1 {
		right := board[i][j+1]
		if !isVisited(i, j+1, oVisited) && 'O' == right {
			p := []int{i, j + 1}
			osAround = append(osAround, p)
			*oVisited = append(*oVisited, p)
		}
	}

	// up
	if i > 0 {
		up := board[i-1][j]
		if !isVisited(i-1, j, oVisited) && 'O' == up {
			p := []int{i - 1, j}
			osAround = append(osAround, p)
			*oVisited = append(*oVisited, p)
		}
	}

	return osAround
}

func findOs(board [][]byte, points [][]int, oVisited *[][]int) [][]int {
	var os [][]int
	for _, p := range points {
		os = append(os, findOsAround(board, p[0], p[1], oVisited)...)
	}
	return os
}

func solve(board [][]byte) {
	if 0 == len(board) {
		return
	}

	var oVisited [][]int

	// visit all 'O's from boundaries that is 'O'
	os := findOBoundaries(board, &oVisited)
	for {
		os = findOs(board, os, &oVisited)
		if len(os) == 0 {
			break
		}
	}

	// alter the not visited 'O' to 'X'
	for i := range board {
		for j := range board[i] {
			if 'O' == board[i][j] && !isVisited(i, j, &oVisited) {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'O', 'X', 'O'},
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
	}
	solve(board)
	fmt.Println(board)
}
