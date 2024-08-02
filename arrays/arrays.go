package arrays

import "fmt"

/*
	Array:
		A sequential data structure that store elements of the same types.
		Normally arrays have fixed size. However, most programming languages
		provides abstractions over array called dynamic arrays, For examples:
		Slices are dynamic arrays in Go or Vector in C++.

	NOTE:
		- Just like in C/C++, Array in Go are zero-indexed. The first
		elements in the array is located at index 0

	Syntax:
		var arrayName [SIZE]dataType

	Declaration and initialization.
	   var arrayName [SIZE]dataType = [SIZE]dataType{data1, data2, data3}

	Accessing Elements in Array:
		var firstElement int = arrayName[0]
		var secondElement int = arrayName[1]
		var lastElement int = arrayName[len(arrayName) - 1]
	Update Element value:
		Update the first element
		arrayName[0] = newValue
*/

// Arrays can also be passed to the function as params.
/*
	Two ways pass array as an arguments to the function.
	1. Pass as a value (essentially passing a copy)
	2. Pass as a pointer
*/
func GetAverage(numbers *[]int, size int) float32 {
	var sum float32

	for _, number := range *numbers {
		sum += float32(number)
	}

	return sum / float32(size)
}

func LearnArrays() {
	const NUMBER_OF_ELEMENTS int = 10
	var elements [NUMBER_OF_ELEMENTS]int

	// Initializing array.
	for i := 0; i < len(elements); i++ {
		elements[i] = i
	}

	total := 0
	// Summing all elements in the array
	for _, element := range elements {
		total += element
	}

	fmt.Printf("the total is %d\n", total)

	// Multidimentional array.
	const ROWS, COLUMNS int = 3, 3

	var matrix [ROWS][COLUMNS]int

	for row := 0; row < ROWS; row++ {
		for column := 0; column < COLUMNS; column++ {
			matrix[row][column] = row * column
		}
	}

	fmt.Println(matrix)

	// Get Average of numbers
	numbers := []int{10, 20, 30}
	fmt.Printf("the average is %f\n", GetAverage(&numbers, len(numbers)))
}
