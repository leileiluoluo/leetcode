package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	maxInterval := (numRows - 1) << 1
	interval := maxInterval
	after := ""
	for i := 0; i < numRows; i++ {
		if numRows-1 == i {
			interval = maxInterval
		}
		for j, no := i, 0; j < len(s); no++ {
			after += string(s[j])
			if i > 0 && i < numRows-1 && 1 == no&1 {
				j += maxInterval - interval
				continue
			}
			j += interval
		}
		interval -= 2
	}
	return after
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
