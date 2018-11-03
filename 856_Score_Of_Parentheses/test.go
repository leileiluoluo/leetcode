package main

import "fmt"

func scoreOfParentheses(s string) int {
	chars := []rune(s)
	r := 0
	for i := 0; i < len(chars)-1; i++ {
		c, next := chars[i], chars[i+1]
		if '(' == c && '(' == next {
			sub := ""
			depth := 1
			for i = i + 1; i < len(chars); i++ {
				if '(' == chars[i] {
					depth++
				} else {
					depth--
				}
				if 0 == depth {
					break
				}
				sub += string(chars[i])
			}
			if len(chars)-1 == i {
				return 2 * scoreOfParentheses(sub)
			}
			return 2*scoreOfParentheses(sub) + scoreOfParentheses(s[i+1:])
		} else if '(' == c && ')' == next {
			sub := ""
			depth := 1
			for i = i + 2; i < len(chars); i++ {
				if '(' == chars[i] {
					depth++
				} else {
					depth--
				}
				if 0 == depth {
					break
				}
				sub += string(chars[i])
			}
			return 1 + scoreOfParentheses(sub)
		}
	}
	return r
}

func main() {
	fmt.Println(scoreOfParentheses("(())()"))
}
