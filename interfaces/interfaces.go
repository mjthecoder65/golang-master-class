package interfaces

import "fmt"

/*
	Interface is a type that specifies a contract of methods without a concrete implementation.
	Interface are named collections of method signatures.

	NOTE:
	 - Go use implicit implementation, meaning you don't need to explicitly declare that a type
		implement  an interface.
	 - Types implement the interface by definiting the required methods.
	 - Empty interface can hold any value.
	 - Type assertions and type switches allow you to interact with underlying values of interfaces.
	 - Interface composition allows you to create complex interfaces from simpler ones.
*/

type Animal interface {
	Speak() string
}

type Dog struct {
	Name  string
	Color string
}

func (dog Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name  string
	Color string
}

func (cat Cat) Speak() string {
	return "Meow!"
}

func PrintSpeak(animal Animal) {
	fmt.Println(animal.Speak())

}

/*
Empty Interface:
Has no method. it can hold any values of any of type because it has not method.
*/
func PrintAnyThing(value interface{}) {
	fmt.Println(value)
}

/*
	Interface composition:
	You can write combine simple interfaces to a more complex one
*/

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func LearnInterfaces() {
	fmt.Println("Hello World!")
}
