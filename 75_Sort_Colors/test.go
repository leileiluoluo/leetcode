package main

import "fmt"

func sortColors(nums []int) {
	swap := func(nums []int, i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}

	start, end := 0, len(nums)-1
	for i := start; i <= end; {
		switch nums[i] {
		case 0:
			if start < i {
				swap(nums, i, start)
			}
			start++
			i++
		case 1:
			i++
		case 2:
			if nums[end] < 2 {
				swap(nums, i, end)
			}
			end--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
