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

type ExchangeAccount struct {
	AccountNumbe int
	Balance      float64
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

func (account *ExchangeAccount) Deposit(amount float64) {
	account.Balance += amount
}

func (account ExchangeAccount) GetCurrentBalance() float64 {
	return account.Balance
}

func (account *ExchangeAccount) Withdraw(amount float64) {
	account.Balance -= amount
}

func (account ExchangeAccount) GetAccountNumber() int {
	return account.AccountNumbe
}

func TestStruct() {

	fmt.Println("Learning about structs.")

	account := ExchangeAccount{AccountNumbe: 12345, Balance: 1000.50}
	fmt.Println(account.AccountNumbe)
	fmt.Println(account.AccountNumbe)
}
