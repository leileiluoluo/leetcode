package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	currLen, maxLen := 0, 0
	preStr := ""
	for _, c := range []rune(s) {
		if strings.Contains(preStr, string(c)) {
			i := strings.LastIndex(preStr, string(c))
			if len(preStr)-1 == i {
				preStr = string(c)
				currLen = 1
			} else {
				preStr = string(preStr[i+1:]) + string(c)
				currLen = len(preStr)
			}
			continue
		}
		preStr += string(c)
		currLen++
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}

func main() {
	s := "ohvhjdml"
	fmt.Println(lengthOfLongestSubstring(s))
}
