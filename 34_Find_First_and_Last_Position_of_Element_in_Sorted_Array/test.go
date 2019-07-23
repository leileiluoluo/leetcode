package main

import "fmt"

func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if target == nums[mid] {
			r := make([]int, 2)
			i := mid
			// find first position
			for i >= 0 && target == nums[i] {
				i--
			}
			i++
			r[0] = i
			// find last position
			for i <= len(nums)-1 && target == nums[i] {
				i++
			}
			r[1] = i - 1
			return r
		}
		if target > nums[mid] {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return []int{-1, -1}
}

func main() {
	nums, target := []int{5, 7, 7, 8, 8}, 8
	fmt.Println(searchRange(nums, target))
}
