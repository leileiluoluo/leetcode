package main

import (
	"fmt"
	"strings"
	"unicode"
)

func calc(chars []rune, start, end int) int {
	r := 0
	pre := 'x'
	for i := start; i < end; i++ {
		c := chars[i]
		if i > start {
			pre = chars[i-1]
		}
		switch c {
		case '(':
			depth := 1
			start = i + 1
			for i = start; i < end && 0 != depth; i++ {
				switch chars[i] {
				case '(':
					depth++
				case ')':
					depth--
				}
			}
			switch pre {
			case 'x', '+':
				r += calc(chars, start, i)
			case '-':
				r -= calc(chars, start, i)
			}
		case '+', '-':
		default:
			v := 0
			for ; i < end && chars[i] >= '0' && chars[i] <= '9'; i++ {
				v = v*10 + int(chars[i]-'0')
			}
			switch pre {
			case 'x', '+':
				r += v
			case '-':
				r -= v
			}
		}
	}
	return r
}

func calculate(s string) int {
	s = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
	return calc([]rune(s), 0, len(s))
}

func main() {
	s := "2-4-(8+2-6+(8+4-(1)+8-10))"
	r := calculate(s)
	fmt.Println(r)
}
