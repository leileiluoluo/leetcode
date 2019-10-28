package main

import "fmt"

func decode(s string, i int, table map[int]int) int {
	if len(s) == i {
		return 1
	}
	if '0' == s[i] {
		return 0
	}
	// if calculated before, return value directly
	if v, ok := table[i]; ok {
		return v
	}
	num := decode(s, i+1, table)
	if i < len(s)-1 {
		if '1' == s[i] {
			num += decode(s, i+2, table)
		} else if '2' == s[i] && s[i+1] <= '6' {
			num += decode(s, i+2, table)
		}
	}
	table[i] = num
	return num
}

func numDecodings(s string) int {
	table := make(map[int]int)
	return decode(s, 0, table)
}

func main() {
	s := "226"
	fmt.Println(numDecodings(s))
}
