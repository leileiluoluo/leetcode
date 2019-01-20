package main

import "fmt"

func numberOfArithmeticSlices(a []int) int {
	if len(a) < 3 {
		return 0
	}

	slices := []int{}
	slice := 2
	preInterval := a[1] - a[0]
	for i := 2; i < len(a); i++ {
		interval := a[i] - a[i-1]
		if interval == preInterval {
			slice++
			if len(a)-1 == i && slice > 2 {
				slices = append(slices, slice)
			}
			continue
		}
		if slice > 2 {
			slices = append(slices, slice)
		}
		slice = 2
		preInterval = interval
	}

	f := func(slice int) int {
		if slice < 3 {
			return 0
		}
		num := 0
		for i := 3; i <= slice; i++ {
			num += (slice - i) + 1
		}
		return num
	}

	sum := 0
	for _, slice := range slices {
		sum += f(slice)
	}
	return sum
}

func main() {
	num := numberOfArithmeticSlices([]int{2, 4, 6, 8, 10})
	fmt.Println(num)
}
