package dsa

import "errors"

type Stack[T any] struct {
	elements []T
}

func (stack *Stack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack[T]) Pop() (T, error) {
	topElement, err := stack.Top()
	if err != nil {
		return topElement, err
	}

	stack.elements = stack.elements[:len(stack.elements)-1]
	return topElement, nil
}

func (stack *Stack[T]) Top() (T, error) {
	if stack.Empty() {
		var zeroValue T
		return zeroValue, errors.New("error: empty stack")
	}

	topIndex := len(stack.elements) - 1
	return stack.elements[topIndex], nil
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.elements) == 0
}
