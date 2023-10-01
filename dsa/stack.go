package dsa

// Stack is a LIFO data structure.
// It has two main operations: push and pop.
// Push: adds an element to the top of the stack, while pop removes the top element from the stack.
// Pop: removes the top element from the stack.
// Top: returns the top element from the stack without removing it.
// IsEmpty: returns true if the stack is empty, otherwise false.
// Size: returns the number of elements in the stack.
// Clear: removes all elements from the stack.
// Stack interface.

type IStack interface {
	Push(element int)
	Pop() int
	Top() int
	IsEmpty() bool
	Size() int
	Clear()
}

type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}

	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Clear() {
	s.elements = make([]int, 0)
}

// CreateStack creates a new stack.
func CreateStack() IStack {
	return &Stack{}
}
