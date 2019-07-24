package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	i, j := 0, 0
	for i < len(intervals) {
		merged := false
		for j = i + 1; j < len(intervals); {
			x, y := intervals[i], intervals[j]
			if min(x[1], y[1]) >= max(x[0], y[0]) {
				merged = true
				// fix intervals[i]
				intervals[i][0], intervals[i][1] = min(x[0], y[0]), max(x[1], y[1])
				// remove intervals[j]
				intervals[j] = intervals[len(intervals)-1]
				intervals = intervals[:len(intervals)-1]
				continue
			}
			j++
		}
		if merged {
			continue
		}
		i++
	}
	return intervals
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
