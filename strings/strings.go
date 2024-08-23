package strings

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

/*
	strings are readonly/immutables slice of bytes.
	Libraries for manipulating strings.

	1. Unicode
	2. regexp
	3. strings
	3. strconv

	A string literal holds a varild UTF-8 sequences called runes.
	A strings holds arbitrary numbers of bytes.

	Rune is a type that represents a Unicode code point.
	It is essentially an alias for the int32 type.

	In Go, rune is used to handle individual characters in strings, and allows you
	to work a wide range of characters beyong the basic ASCII, including symbols from various
	languages and script. A basic ASCII can be reprented by byte type (8bits).

	byte is an alias for uint8

	Note:
		1. No concepts of characters like in C++.
		2. Strings are immutable.
			var name string = "Michael"
			name[0] = rune('2') Does not work.
			- However you can modify by converting string to a slice of rune
			Modify the charater and convert it back to string
			runes := []rune(name)
			runes[0] = 'H'
			name = string(runes)
*/

func MasteringStrings() {
	var greeting string = "Hello World"
	fmt.Println(greeting)
	fmt.Println(reflect.TypeOf(greeting[0]))

	// for index, code := range greeting {
	// 	fmt.Printf("Type : %s, index: %d, byte: %d\n", reflect.TypeOf(code), index, code)

	// 	fmt.Println(string(code))
	// 	var charRune rune = rune(code)

	// 	fmt.Println(string(charRune)) // Convert rune to string

	// 	charCode := rune(greeting[index]) // converting to tune.
	// 	fmt.Println(charCode)
	// }

	// Getting length of the string
	fmt.Printf("Length of the string: %d\n", len(greeting))

	// Getting substring.
	const (
		START int = 0
		END   int = 5
	)

	// String slicing.
	fmt.Printf("Substring: %s\n", greeting[START:END])

	var name string = "Michael Jordan"

	// Check if container a subtring
	var substr string = "Michael"
	fmt.Println(strings.Contains(name, substr))

	// Return index of the first occurency of a substring. -1 if not found.
	fmt.Println(strings.Index(name, substr))

	// Replace occurency of a substring with another substr.
	var toBeReplaced string = "Jordan"
	var newSubtr = "Ngowi"
	fmt.Println(strings.Replace(name, toBeReplaced, newSubtr, 1))

	// Splitting string.
	names := strings.Split(name, " ")

	for _, name := range names {
		fmt.Println(name)
	}

	firstname := "Michael"
	lastname := "Ngowi"

	// Concatenate string.
	fullname := firstname + " " + lastname

	fmt.Println(fullname)

	// Join a slice of strings into a single string.
	fullname = strings.Join([]string{firstname, lastname}, " ")
	fmt.Println(fullname)

	// Upper and Lower
	fmt.Println(strings.ToUpper(firstname))
	fmt.Println(strings.ToLower(lastname))

	// Trim leadning and trailing whitespace
	fmt.Println(strings.Trim(" username@gmail.com ", " "))

	// Suffix and Prefix check
	fmt.Println(strings.HasPrefix(firstname, "Mi"))
	fmt.Println(strings.HasSuffix(firstname, "el"))

	// String to number
	number, err := strconv.Atoi("1235")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The converted number is %d\n", number)

	// Integer to String
	fmt.Println(strconv.Itoa(1500))

	// Parse Fuctions
	/*
		strconv.ParseBool(str string) (bool, error)
		strconv.ParseFloat(s string, bitSize int) (float64, error)
		strconv.ParseInt(s string, base int, bitSize int) (int64, error)
		strconv.ParseUint(s string, base int, bitSize int) (uint64, error)

		b, err := strconv.ParseBool("true")
		f, err := strconv.ParseFloat("3.14", 64)
		i, err := strconv.ParseInt("123", 10, 64)
		u, err := strconv.ParseUint("123", 10, 64)


	*/

	// Format Function.
	/*
		s := strconv.FormatBool(true)                // "true"
		s := strconv.FormatFloat(3.14, 'f', -1, 64)  // "3.14"
		s := strconv.FormatInt(123, 10)              // "123"
		s := strconv.FormatUint(123, 10)             // "123"


		strconv.FormatBool(b bool) string
		strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string
		strconv.FormatInt(i int64, base int) string
		strconv.FormatUint(i uint64, base int) string /
	*/
}
