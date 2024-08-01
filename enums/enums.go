package enums

import "fmt"

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func Enums() {
	fmt.Println("Learning about Enums in Go!")

	var today Day = Wednesday

	fmt.Println(today)

}
