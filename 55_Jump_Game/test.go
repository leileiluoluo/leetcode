package main

import "fmt"

func canJump(nums []int) bool {
	reach := 0
	for i, num := range nums {
		if reach >= len(nums)-1 {
			return true
		}
		if 0 == num && reach <= i {
			return false
		}
		if i+num > reach {
			reach = i + num
		}
	}
	return false
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
