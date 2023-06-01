package main

import "fmt"

func testMaps() {
	// Creating a map
	// A map is a collection of key-value pairs.
	// A map is similar to a dictionary in Python.
	// A map is also similar to a hash table in other languages.
	// A map is an unordered collection.
	// A map is a reference type.
	// A map is created using the make() function.
	// The make() function takes in the data type of the key and the value.
	// The make() function also takes in the initial capacity of the map.
	// The initial capacity is optional.
	// The initial capacity is the number of key-value pairs that the map can hold.
	// The initial capacity is used to improve performance.

	// Syntax:
	// mapName := make(map[keyDataType]valueDataType, initialCapacity)

	// Example:
	balances := make(map[string]int, 10)

	// Adding key-value pairs to a map
	balances["Michael"] = 100
	balances["Jerry"] = 200

	// Deleting a key from a map
	// Syntax: delete(mapVariable, key)
	delete(balances, "Jerry")

	// Check if key exists in a map
	// Syntax:
	// value, ok := mapName[key]
	// ok is a boolean value
	// ok is true if the key exists in the map

	value, exists := balances["Michael"]

	if exists {
		fmt.Printf("Michael's balance is %v\n", value)
	} else {
		fmt.Println("Key Michael does not exist in the map")
	}

	// Interating over a map
	for key, value := range balances {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	// Maps are unordered data structures. The order of the key-value pairs is not guaranteed.
	// If you need to maintain the order of the key-value pairs, use a slice of structs instead.
	// If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.

}
