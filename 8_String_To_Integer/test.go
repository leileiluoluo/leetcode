package main

import "fmt"

const (
	MaxInt = 1<<31 - 1
	MinInt = -(1 << 31)
)

func myAtoi(str string) int {
	v := 0
	i := 0
	negtive := false

	// trim blank prefix
	for ; i < len(str); i++ {
		if ' ' != str[i] {
			break
		}
	}
	if i > len(str)-1 {
		return v
	}

	// first non-blank char
	if '+' == str[i] {
		i++
	} else if '-' == str[i] {
		negtive = true
		i++
	} else if str[i] < '0' || str[i] > '9' {
		return v
	}

	// integer
	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		v = 10*v + int(str[i]-'0')
		if negtive {
			if -v <= MinInt {
				return MinInt
			}
			continue
		}
		if v < 0 || v > MaxInt {
			return MaxInt
		}
	}

	if negtive {
		v = -v
	}
	return v
}

func main() {
	str := "2147483648"
	fmt.Println(myAtoi(str))
}
