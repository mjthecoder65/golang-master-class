package loops

import "fmt"

type Student struct {
	Name string
	Age  int
}

func LearnLoops() {
	/*
		Different kinds of for loops in Golang
		1. Traditional for init; condition; increment
		2. Range loop: for index, numbers := range numbers
		3. While loop format: for condition

		for (condition) | (init; condition; increment) | Range
	*/
	/*
	 Check the following example;
	*/

	numbers := []int{11, 13, 10, 20, 50, 250}

	i := 0 // init

	// Using for (condition): This is similar to while loop
	fmt.Println("Printing numbers using: for (condition)")
	for i < len(numbers) {
		fmt.Println(numbers[i])
		i++ // increment
	}

	// Using for (init; condition; increment)
	fmt.Println("Printing numbers using : for (init; condition; increment)")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %d and Value: %d\n", i, numbers[i])
	}

	// Printing using for (Range)
	fmt.Println("Printing numbers using : for (Range) loop")
	for index, value := range numbers {
		fmt.Printf("Index: %d and Value : %d\n", index, value)
	}

	students := make(map[int]Student)
	students[1] = Student{Name: "Michael", Age: 31}
	students[2] = Student{Name: "Jerry", Age: 29}

	for key, value := range students {
		fmt.Println(key, value)
	}

	// Break and Continue Keywords.
	/*
		break: used to stop the loop and pass control to statement right next after the loop.
		continue: This move to the next iteration, ignoring all the statements after its.
	*/

	// Example: using break
	fmt.Println("Print all numbers excluding and stop the loop when encountering even number")
	for _, number := range numbers {
		fmt.Println(number)
		if number%2 == 0 {
			break
		}
	}

	// Example: using break
	fmt.Println("Print all numbers excluding odd numbers ")
	for _, number := range numbers {
		if number%2 == 1 {
			continue
		}
		fmt.Println(number)
	}
}
