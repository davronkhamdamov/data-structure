package main

import "fmt"

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}
func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(12)
	myStack.Push(53)
	myStack.Push(41)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())

}
