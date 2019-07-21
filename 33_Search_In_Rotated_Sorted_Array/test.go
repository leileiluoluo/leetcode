package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[start] <= nums[end] { // eg. [0, 1, 2, 3, 4, 5, 6, 7]
			if target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { // eg. [4, 5, 6, 7, 0, 1, 2]
			if target >= nums[start] {
				if nums[mid] >= nums[start] &&
					target > nums[mid] {
					start = mid + 1
				} else {
					end = mid - 1
				}
			} else {
				if nums[mid] >= nums[start] {
					start = mid + 1
				} else {
					if target < nums[mid] {
						start++
						end = mid - 1
					} else {
						start = mid + 1
					}
				}
			}
		}
	}
	return -1
}

func main() {
	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 2
	fmt.Println(search(nums, target))
}
