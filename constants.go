package main

import "fmt"

// constant is a name of or identifier for a fixed value.
// Constants are immutable.
// Keyword const is used to declare a constant.
// You must assign a value to a constant when you declare it, you cannot assign later as with variables.

// Syntax: const CONSTANT_NAME dataType = value
// Example: const PI float64 = 3.14159

// Naming convention for constants:
// 1. Use all upper case letters.
// 2. Use underscore to separate words.

func TestingConstants() {
	// Declaring a constant
	const NUMBER_OF_ELEMENTS int = 10

	var products [NUMBER_OF_ELEMENTS]string
	products[0] = "Computer"

	// Declaring multiple constants using parentheses.
	const (
		PI       float64 = 3.14159
		LANGUAGE string  = "Go"
	)

	fmt.Printf("The value of PI is %f\n", PI)
	fmt.Printf("The value of LANGUAGE is %s\n", LANGUAGE)
	fmt.Println(products)

}
