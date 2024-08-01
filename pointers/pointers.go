package pointers

import "fmt"

/*
Remember:
	1. Every variable is a memory location
	2. Every Memory location has its address. In Go, memory address can be accessed by
		using & (ampersand operator - unary operator)
Synatax:
	&variable_name
	memory_address := &variable_name
*/

func Pointers() {
	var age int = 20
	fmt.Printf("Address of a variable age is : %x\n", &age)
}
