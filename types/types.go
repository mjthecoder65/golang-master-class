package main

import (
	"fmt"
	"reflect"
)

// Checking types in Golang

func CheckingTypes() {
	var number int = 10
	var name string = "Jane"
	var isTrue bool = true
	var floatNumber float64 = 10.50
	var elements []int = []int{1, 2, 3}
	var student map[string]string = map[string]string{"name": "Jane", "age": "20"}
	var ptr *int = &number
	var function func() int

	// Checking types\
	fmt.Println(reflect.TypeOf(number))
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(isTrue))
	fmt.Println(reflect.TypeOf(floatNumber))
	fmt.Println(reflect.TypeOf(elements))
	fmt.Println(reflect.TypeOf(student))
	fmt.Println(reflect.TypeOf(ptr))
	fmt.Println(reflect.TypeOf(function))

}

func LearnTypes() {
	// Type casting
}
