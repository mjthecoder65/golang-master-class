package main

import "fmt"

const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func Enums() {
	fmt.Println("Hello, World! Learning about Enums in Go today.")
	var today int = Wednesday

	fmt.Println(today)
}
