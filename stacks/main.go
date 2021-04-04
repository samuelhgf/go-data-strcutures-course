package main

import (
	"errors"
	"fmt"
)

// Stack - a stack data structure
type Stack struct {
	Elements []int
}

// Pop - pops out the last element from the stack and updated it
func (s *Stack) Pop() (int, error) {
	var lastElement int

	// Verify if stack is empty
	if s.IsEmpty() {
		return 0, errors.New("empty stack")
	}

	// Pop the last element and update the Stack
	lastElement, s.Elements = s.Elements[s.Length() - 1], s.Elements[0:s.Length() - 1]

	// Return the last element popped from the Stack
	return lastElement, nil
}

// Push - add an element to the top of the Stack
func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

// IsEmpty - returns true case the Stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

// Length - returns the length of the Stack
func (s *Stack) Length() int {
	return len(s.Elements)
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(3)
	stack.Push(5)
	stack.Push(7)

	fmt.Println(stack)

	elem, _ := stack.Pop()

	fmt.Println(elem)
	fmt.Println(stack)

	stack.Push(10)

	fmt.Println(stack)

	elem, _ = stack.Pop()

	fmt.Println(elem)
	fmt.Println(stack)
}
