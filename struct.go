package main

import "fmt"

// Struct is used to create a compound data type.
// It is used to group related data together.
// It is similar to a class in OOP.
// Data fields in a struct can have different data types.

// Syntax:
/*
type structName struct {
	field dataType
	field dataType
}
*/

type Account struct {
	accountNumber int
	balance       float64
}

// Struct methods.
// A struct method is a function that is associated with a struct.
// A struct method is similar to a method in OOP.

// Syntax:
/*
func (structName structType) methodName() {
	// Code goes here
}
*/

func (account *Account) Deposit(amount float64) {
	account.balance += amount
}

func (account Account) GetCurrentBalance() float64 {
	return account.balance
}

func (account *Account) Withdraw(amount float64) {
	account.balance -= amount
}

func (account Account) GetAccountNumber() int {
	return account.accountNumber
}

func TestStruct() {

	fmt.Println("Learning about structs.")

	account := Account{accountNumber: 12345, balance: 1000.50}
	fmt.Println(account.accountNumber)
	fmt.Println(account.accountNumber)
}
