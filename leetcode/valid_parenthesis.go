package leetcode

// URL : https://leetcode.com/problems/valid-parentheses/description/
// Title : Valid Parentheses

import (
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() (rune, error) {
	if s.IsEmpty() {
		return '0', fmt.Errorf("stack is empty")
	}

	lastIndex := len(s.items) - 1
	removedItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return removedItem, nil
}

func (s *Stack) Peek() (rune, error) {
	if s.IsEmpty() {
		return '0', fmt.Errorf("stack is empty")
	}
	lastIndex := len(s.items) - 1
	return s.items[lastIndex], nil
}

func IsValidParenthesis(s string) bool {
	stack := Stack{}

	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			// Pushing an element to the stack.
			stack.Push(char)
		} else if char == ')' || char == '}' || char == ']' {
			// Returns false if char is close parenthesis but the stack is empty
			if stack.IsEmpty() {
				return false
			}

			topElement, _ := stack.Peek()

			stack.Pop()
			// Check if top element is equal to the cl
			if topElement != mapping[char] {
				return false
			}

		}
	}
	return stack.IsEmpty()
}
