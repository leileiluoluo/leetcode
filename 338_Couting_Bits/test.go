func countBinaryOneInHexUnit(n int) int {
	countOne := 0
	switch n {
	case 0:
		countOne = 0
	case 1, 2, 4, 8:
		countOne = 1
	case 3, 5, 6, 9, 10, 12:
		countOne = 2
	case 7, 11, 13, 14:
		countOne = 3
	case 15:
		countOne = 4
	}
	return countOne
}

func countBinaryOne(n int) int {
	// remain
	r := 0
	countOne := 0
	for n > 0 {
		n, r = n>>4, n%16
		countOne += countBinaryOneInHexUnit(r)
	}
	return countOne
}

func countBits(num int) []int {
	s := make([]int, num+1)
	for i := 0; i <= num; i++ {
		s[i] = countBinaryOne(i)
	}
	return s
}
