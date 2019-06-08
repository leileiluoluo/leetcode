package main

import "fmt"

type stack struct {
	stores []int
}

func (s *stack) top() int {
	return s.stores[len(s.stores)-1]
}

func (s *stack) pop() {
	s.stores = s.stores[:len(s.stores)-1]
}

func (s *stack) push(e int) {
	s.stores = append(s.stores, e)
}

func (s *stack) len() int {
	return len(s.stores)
}

func validateStackSequences(pushed []int, popped []int) bool {
	s := new(stack)
	i, j := 0, 0
	for ; i < len(popped); i++ {
		pop := popped[i]
		if s.len() > 0 && s.top() == pop {
			s.pop()
			continue
		}
		if j == len(pushed) {
			return false
		}
		for j < len(pushed) {
			e := pushed[j]
			if pop != e {
				s.push(e)
				j++
				continue
			}
			j++
			break
		}
	}
	return 0 == s.len() && len(popped) == i
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
}
