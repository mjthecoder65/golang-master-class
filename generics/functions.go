package generics

import "fmt"

func Swap[T any](a, b T) (T, T) {
	return b, a
}

// This is a way to define a type set.
// This was also introduced in Go 1.18 as a part of
// Generic feature.
type Number interface {
	int | int64 | float32 | float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func LearnFunctionGenerics() {
	fmt.Println("Learning type generics in Go.")
}
