package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	charCounts := make(map[rune]int)
	for _, c := range s {
		charCounts[c] += 1
	}

	var counts []int
	countChars := make(map[int][]rune)
	for c, count := range charCounts {
		if chars, ok := countChars[count]; ok {
			countChars[count] = append(chars, c)
		} else {
			countChars[count] = []rune{c}
			counts = append(counts, count)
		}
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	var chars []rune
	for _, count := range counts {
		for _, c := range countChars[count] {
			i := count
			for i > 0 {
				chars = append(chars, c)
				i--
			}
		}
	}
	return string(chars)
}

func main() {
	s := "aAbb"
	fmt.Println(frequencySort(s))
}
