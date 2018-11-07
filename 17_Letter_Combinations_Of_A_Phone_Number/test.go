package main

import "fmt"

var (
	mapping = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	var combinations []string
	if 0 == len(digits) {
		return combinations
	}
	if 1 == len(digits) {
		return mapping[digits[0]]
	}
	for _, letter := range mapping[digits[0]] {
		for _, suffix := range letterCombinations(digits[1:]) {
			combination := string(letter) + suffix
			combinations = append(combinations, combination)
		}
	}
	return combinations
}

func main() {
	fmt.Println(letterCombinations("234"))
}
