package structs

import (
	"fmt"
	"reflect"
)

/*
	Structures: User-defined composite data types that can hold data of different types.
	Structure as used to represent a record - a collection of data of different types as a unit.
	Example:
	You can create  struct of called book with the following attributes:
		Title
		Author
		Subject
		Book ID

	The structures attributes are also called members.

	To define a structure, you must use both keywords : type and struct
	Syntax:

	type Book struct {
		member_1 data_type
		member_2 data_type
		...
		member_N data_type
	}

	Access:
		Member accessed using member access operator (.)
		Example: Book.Name

	Structural Method:
		Allows developers to create function that operates on members of the structure.
		Methods have a special receiver arguments. Receiver arguments appears in its own
		list beween func keyword and the method name.

	Value vs Pointer Receivers:
		Value Receiver: A copy of the struct is passed to the method. Modifications made within the method
		do not affect the original struct.
		Pointer Receiver: A pointer to the struct is passed to the method. Modification made with the medhod
		affects the original struct.
*/

type Book struct {
	Title   string
	Author  string
	Subject string
	BookId  int
}

func (book Book) Print() {
	fmt.Printf("Titlle: %s\n", book.Title)
	fmt.Printf("Author: %s\n", book.Author)
	fmt.Printf("Subject: %s\n", book.Subject)
	fmt.Printf("BookId: %d\n", book.BookId)
}

/*
Passing structure as an argument.
*/

func PrintBookPassedByValue(book Book) {
	fmt.Printf("Title: %s\n", book.Title)
	fmt.Printf("Author: %s\n", book.Author)
	fmt.Printf("Subject: %s\n", book.Subject)
	fmt.Printf("BookId: %d\n", book.BookId)
}

func PrintBookPassedByPointer(book *Book) {
	fmt.Printf("Title: %s\n", book.Title)
	fmt.Printf("Author: %s\n", book.Author)
	fmt.Printf("Subject: %s\n", book.Subject)
	fmt.Printf("BookId: %d\n", book.BookId)
}

// You can as well create Nested Structs.
type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

/*
Struct Tags:
Annotations added to struct fields that provide metata. This metadata are used
for purposes like encoding and decoding  data(JSON, XML), database interaction, or other forms
of data serialization.

Common Use Case:

	JSON serialization.
*/

// Example:
type Token struct {
	Price string `json:"price"` // specificies that Name should be encoded as "name" in JSON
	Name  int    `json:"age"`   // Spacifies that Age should be encoded as "age"
}

type Investment struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	ProductName string  `gorm:"column:person_name" json:"product_name"`
	Amount      float64 `gorm:"column:amount" json:"amount"`
	Currency    string  `gorm:"column:currency" json:"currency"`
}

// Embbeded struct

type HomeAddress struct {
	Street string
	City   string
}

type User struct {
	Name        string
	Age         int
	HomeAddress // Embedded struct
}

func LearnStructures() {
	var book Book = Book{
		Title:   "Deep Work",
		Author:  "Cal NewPort",
		Subject: "Self-development",
		BookId:  2, // Remember to add comma after the last member
	}

	// book.Title = "Deep Work"
	// book.Author = "Cal NewPort"
	// book.Subject = "Self Development"
	// book.BookId = 2

	PrintBookPassedByValue(book)

	/*
		A pointer to a structure. You can create a pointer to a structure.
		Member Access: A way to access members doesn't change even when using pointer.
		In C++, you use . for member access but -> when you have a pointer to a struct.
	*/

	var bookPtr *Book = &book

	PrintBookPassedByPointer(bookPtr)

	// Modify book id using pointer.
	bookPtr.BookId = 1

	// Check if the id of the book has been changed.
	PrintBookPassedByValue(book)

	// Printing the book information using a Method.
	book.Print()

	// Creating annoymous struct
	point := struct {
		X int
		Y int
	}{X: 10, Y: 20}

	fmt.Println(point)

	// Accessing Tags Using reflect method.
	investment := Investment{
		ID:          1,
		ProductName: "BTC Income",
		Amount:      2.5893,
		Currency:    "BTC",
	}

	t := reflect.TypeOf(investment)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s : %s\n", field.Name, field.Tag.Get("json"))
	}

	// Embedded struct and member access.
	user := User{
		Name: "Michael",
		Age:  30,
		HomeAddress: HomeAddress{
			Street: "Jongam-ro",
			City:   "Seoul",
		},
	}

	fmt.Println(user.Street)
	fmt.Println(user.Age)
	fmt.Println(user.Name)
}
