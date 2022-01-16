package main

import "fmt"

/*
	Stack is a linear data structure which follows a particular order in which operations are performed.
	The orders are LIFO (Last in, First out).
	2 operations can be performed on the stack, and that is the pop and the push.
*/

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items) - 1]
}

func main() {
	implementStack := Stack{}
	implementStack.Push(1)
	implementStack.Push(2)
	implementStack.Push(3)
	fmt.Println(implementStack.items)
	implementStack.Pop()
	implementStack.Pop()
	implementStack.Pop()
	fmt.Println(implementStack)
}