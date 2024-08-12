package controls

import "fmt"

/*
	if statment : A control structure that allows you to execute
	certain block of codes on a condition.
	If the condition evaluates to true, the code block is with the if statement
	is executed, otherwise skipped.

	Syntanx:
		if condition1 {
			// Block to run when condition1 is true.
		} else if condition2 {
		// Block to run when condition2 is true
		} else {
		// Block run when all conditions are false.
		}

	if statments can be nested as well.

	if condition 1 {
	    if condition 2 {
			// Block of code: Executed when both condition 1 and condition 2 are true.
		} else {
		 // Executed when only condition1 is true but condition2 is false.
		}
	}

	Short Variable Declaration in if statment

	if v := expression: condition {
		// here v is only available within if statement.
	}

	IMPOTANT POINTS:
		1. No parenthesis are required to surround the condition. Unlike other languages like C and C++
		2. Braces are mandatory even for single line.
		3. Variable declared in the short variable declaration within an if statments are scopped to if, else if and else.
*/

func learnControlStructure() {
	//
	fmt.Println("Learning Control Structures")
}

func ControlStructures() {
	//

	if age := 10; age < 20 {
		fmt.Printf("Your name is %s, and you are %d years old\n", name, age)
	}
}
