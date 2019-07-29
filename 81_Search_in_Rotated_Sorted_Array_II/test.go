package main

import "fmt"

func search(nums []int, target int) bool {
	begin, end := 0, len(nums)-1
	for begin <= end {
		// trim begin nums of begin == end
		if nums[begin] == nums[end] {
			if target == nums[begin] {
				return true
			}
			for begin < end && nums[begin] == nums[end] {
				begin++
				continue
			}
		}

		mid := (begin + end) / 2
		if target == nums[mid] {
			return true
		}

		// all is asending
		if nums[begin] < nums[end] {
			if target > nums[mid] {
				begin = mid + 1
			} else {
				end = mid - 1
			}
			continue
		}

		// rotated asending
		if target > nums[mid] {
			if nums[mid] >= nums[begin] {
				begin = mid + 1
			} else {
				if target >= nums[begin] {
					end = mid - 1
				} else {
					begin = mid + 1
				}
			}
		} else {
			if nums[mid] >= nums[begin] {
				if target >= nums[begin] {
					end = mid - 1
				} else {
					begin = mid + 1
				}
			} else {
				end = mid - 1
			}
		}
	}
	return false
}

func main() {
	nums, target := []int{3, 1, 1}, 3
	fmt.Println(search(nums, target))
}
