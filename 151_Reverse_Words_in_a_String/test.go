package main

import "fmt"

func reverseWords(s string) string {
	reversed := ""
	lastIndex := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if ' ' == s[i] {
			if i == lastIndex {
				lastIndex--
			} else {
				reversed += s[i+1:lastIndex+1] + " "
				lastIndex = i - 1
			}
		} else {
			if 0 == i && i <= lastIndex {
				reversed += s[i : lastIndex+1]
			}
		}
	}
	if len(reversed) > 1 && ' ' == reversed[len(reversed)-1] {
		reversed = reversed[:len(reversed)-1]
	}
	return reversed
}

func main() {
	s := "  hello world!  "
	fmt.Println(reverseWords(s))
}
