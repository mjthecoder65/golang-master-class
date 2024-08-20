package functions

import "fmt"

func Factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * Factorial(number)
}

// Defininig a custom data type
type Operation func(int, int) int

/*
Type is used create alias for existing type or new custom data types.
such as function types, structures, and interfaces
*/
type Integer int

func AddNums(a int, b int) (total int) {
	total = a + b
	return
}

func Recursion() {
	var fib func(n int) int

	/*
		Functions are first-class in Go.
			That means functions are treated as any value or variables in Go,
			allowing you to assign them to a variable, pass them as argument, and return them from other
			functions.

			Function Type is defined by its signature.Which includes func keyword, parameter
			types and then return types.
	*/

	/*
		Higher Order functions:
			Functions that take other functions as argument or return them as result.
			These are common in go, especially in libraries and frameworks that need to perform
			operations like callback, decorator and events handling.

	*/
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	var op Operation = Add

	result := op(3, 4)

	fmt.Println(result)

	var number Integer = 20

	fmt.Println(number)
}
