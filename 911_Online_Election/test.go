package main

import "fmt"

type TopVotedCandidate struct {
	times    []int
	leadings map[int]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var orders [][]int
	leadings := make(map[int]int)
	for i := 0; i < len(times); i++ {
		t := times[i]
		p := persons[i]
		count := 1
		// orders
		for j, order := range orders {
			if p == order[0] {
				count = order[1] + 1
				orders = append(orders[:j], orders[j+1:]...)
			}
		}
		orders = append(orders, []int{p, count})

		// leadings map
		max := 1
		for i := len(orders) - 1; i >= 0; i-- {
			count := orders[i][1]
			if count > max {
				p = orders[i][0]
				max = count
			}
		}
		leadings[t] = p
	}
	return TopVotedCandidate{times, leadings}
}

func (this *TopVotedCandidate) Q(t int) int {
	// binary search
	q := this.times[0]
	start, end := 0, len(this.times)-1
	for start <= end {
		mid := (start + end) / 2
		if t < this.times[mid] {
			end = mid - 1
			continue
		}
		q = this.times[mid]
		start = mid + 1
	}
	return this.leadings[q]
}

func main() {
	persons := []int{0, 0, 1, 1, 2}
	times := []int{0, 67, 69, 74, 87}
	queries := []int{4, 62, 100, 88, 70, 73, 22, 75, 29, 10}
	obj := Constructor(persons, times)
	for _, q := range queries {
		leading := obj.Q(q)
		fmt.Println(leading)
	}
}
