package main

import "fmt"

func comb(begin, end, k int) [][]int {
	var r [][]int
	for i := begin; i <= end; i++ {
		if 1 == k {
			r = append(r, []int{i})
			continue
		}
		suf := comb(i+1, end, k-1)
		for _, j := range suf {
			r = append(r, append([]int{i}, j...))
		}
	}
	return r
}

func combine(n int, k int) [][]int {
	return comb(1, n, k)
}

func main() {
	fmt.Println(combine(6, 4))
}
