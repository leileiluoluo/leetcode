package main

import "fmt"

func recusiveRob(nums []int, len int, table map[int]int) int {
	if 0 == len {
		return 0
	}
	if 1 == len {
		return nums[0]
	}
	if v, ok := table[len]; ok {
		return v
	}
	max := recusiveRob(nums, len-2, table) + nums[len-1]
	pre := recusiveRob(nums, len-1, table)
	if pre > max {
		max = pre
	}
	table[len] = max
	return max
}

func rob(nums []int) int {
	table := make(map[int]int)
	return recusiveRob(nums, len(nums), table)
}

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}
