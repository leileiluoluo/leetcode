package main

import (
	"fmt"
	"strconv"
)

func factorial(i int) int {
	f := 1
	for ; i >= 1; i-- {
		f *= i
	}
	return f
}

func getPermutaion(s string, k int) string {
	i := len(s)
	if 1 == i {
		return s
	}
	factorial := factorial(i)
	nextFactorial := factorial / i
	if k <= nextFactorial {
		return s[:1] + getPermutaion(s[1:], k)
	}
	c, k := (k-1)/nextFactorial, (k-1)%nextFactorial+1
	if c > 0 {
		s = string(s[c]) + s[:c] + s[c+1:]
	}
	return getPermutaion(s, k)
}

func getPermutation(n int, k int) string {
	s := ""
	for i := 1; i <= n; i++ {
		s += strconv.Itoa(i)
	}
	return getPermutaion(s, k)
}

func main() {
	fmt.Println(getPermutation(4, 24))
}
