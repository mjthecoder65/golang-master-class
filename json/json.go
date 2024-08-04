package json

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Serialization (Marshal):
		A process of converting complex data structure (like objects, or lists)
		into a format that can be stored or transmitted. This format is often
		a string of text or a byte stream. Common serialization format include JSON,
		XML, and Protocals Buffers.

	Deserialization (UnMarshal):
		A reverse process of converting serialized representation back into
		its orginal data structure. Think of it as reconstruction the original object from
		the stored data.

	Why is this important:
		1. Data storage: Serialization allows you to save objects to files or databases for later use.
		2. Data Transmission: Your send serialized data over the networks efficiently as it is often smaller and easier
		to transport than the original object.
		3. Data Sharing: Serialized data can be shared between different systems or programming languages.
*/

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (person Person) Print() {
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
}

func LearnJSON() {
	person := Person{
		Name: "Michael",
		Age:  30,
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(person) // Return a stream of bytes

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	var person2 Person

	err = json.Unmarshal(jsonData, &person2)

	if err != nil {
		log.Fatal(err)
	}

	person2.Print()

}
