package main

import "fmt"

var table = make(map[int]int)

func numTrees(n int) int {
	if 0 == n || 1 == n {
		return 1
	}
	if v, ok := table[n]; ok {
		return v
	}
	num := 0
	for i := 1; i <= n; i++ {
		num += numTrees(i-1) * numTrees(n-i)
	}
	table[n] = num
	return num
}

func main() {
	fmt.Println(numTrees(5))
}
