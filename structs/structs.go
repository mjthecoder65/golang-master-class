package structs

import "fmt"

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
}
