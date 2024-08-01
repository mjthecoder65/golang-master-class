package leetcode

// URL: https://leetcode.com/problems/backspace-string-compare/
// Title: Backspace String Compare

// Return true if slice1 and slice2
// We can use two finger method or we could

func areRuneSlicesEqual(slice1 []rune, slice2 []rune) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for index := range slice1 {
		if slice1[index] != slice2[index] {
			return false
		}
	}
	return true
}

func processBackSpace(str string) []rune {
	stack := []rune{}

	for _, char := range str {
		if char == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, char)
		}
	}
	return stack
}

func backspaceCompare(s string, t string) bool {
	return areRuneSlicesEqual(processBackSpace(s), processBackSpace(t))
}
