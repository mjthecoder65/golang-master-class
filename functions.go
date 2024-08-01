package main

import "fmt"

// Functions are the building blocks of a Go program.
// Functions are used to organize code.
// Functions are used to make code more readable.
// Functions are used to make code reusable.
// Functions are first-class in go.

func TestFunctions() {
	fmt.Println("Learning about functions in go.")

	total := Sum(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)

	numbers := []int{1, 2, 3, 4, 5}

	total = Sum(numbers...)
	fmt.Println("Total:", total)

	double := Multiplier(2) // double is a function that takes an int and returns an int
	fmt.Println(double(5))

}

// This function neither takes paramters nor return value
func SayHello() {
	fmt.Println("Hello")
}

// Takes two arguments but returns no value.
func SendEmail(email string, message string) {
	fmt.Println("Sending email to", email)
	fmt.Println("Message:", message)
}

// Takes two arguments and returns a string.
func GetFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

// In Go you can name the return value of a function.
func Add(x int, y int) (total int) {
	total = x + y
	return // Naked return
}

// The above function is the same as this one:
func Add2(x, y int) (total int) {
	total = x + y
	return total
}

// Multiple return values
func Add3(x, y int) (total int, product int) {
	total = x + y
	product = x * y
	return
}

// Recursion is supported in Go
func Fact(number int) int {
	if number < 1 {
		return 1
	}

	return number * Fact(number-1)
}

// Variadic functions
// A variadic function is a function that takes in a variable number of arguments.
// A variadic function is a function that takes in a slice of arguments.
// fmt.Println() is an example of a variadic function.
func Sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Closure is a function that returns a function.
// The returned function has access to the variables in the outer function.
// The returned function can access and modify the variables in the outer function even after the outer function has finished executing.
// A real life use of closure is in implementation of callback fucntion or even handlers.
// Closure allow you to create functions that can access and modify variable defined outside of their own scope.

func Multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func Sequence() func() int {
	counter := 0

	return func() int {
		counter += 1
		return counter
	}

}
