package functions

import "fmt"

// Functions are the building blocks of a Go program.
// Functions are used to organize code.
// Functions are used to make code more readable.
// Functions are used to make code reusable.
// Functions are used to make code more maintainable.
// Functions are first-class in go.

// Variadic functions
// Variadic functions are functions that take a variable number of arguments.
// Here a function take any number of postional arguments.
func AddNumbers(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Syntax:
/*
func functionName(arg ...type) returnType {
	// code
}
*/
// The three dots (...) are called the ellipsis operator.

func Sum(numbers ...int) int {
	var total int

	for _, number := range numbers {
		total += number
	}

	return total
}

func Println(words ...string) {
	for _, word := range words {
		fmt.Println(word)
	}
}
