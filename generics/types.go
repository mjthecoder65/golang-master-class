package generics

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (stack *Stack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack[T]) Pop() T {
	if stack.Empty() {
		var zero T
		return zero
	}

	element := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return element
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.elements) == 0
}

func LearnTypesGenerics() {
	fmt.Println("Learning type generics in Go.")
}
