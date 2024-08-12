package testing

import (
	"net/http"
	"testing"
)

/*
	Go has a built-in testing framework called testing,
	which makes it straight forward to write tests for your code.

	NOTE: In Go, tests are usually placed in the same package as the codes
	they are testing, and named with suffix _test.go

	Types of testing:
		1. Unit Testing
			Unit tests involves testing individual units or components of a software
			to verify tht each unit functions correctly. In Go, a unit test typically focuses on a single function
			or method.

		2. Integration Testing
			Focuses on testing the interaction between different component of a system.
			These tests ensure that the components work together as expected.

		3. End to End Testing()
			End-to-end testing involves testing the entire application flow, from start to finish,
			to ensure everything works together in a real-world scenario. Usually involves testing the application
			through its user interface.

	Three Concepts in Testing
		1. Fixtures
		2. Mocking
		3. Patching
*/

func Add(a, b int) int {
	return a + b
}

func FetchData(url string) (*http.Response, error) {
	return http.Get(url)
}

func ProcessData(res *http.Response) string {
	return "processed data"
}

func TestAndProcessData(t *testing.T) {
	url := "example.com"
	res, err := FetchData(url)

	if err != nil {
		t.Fatalf("Failed to fetch data:  %v", err)
	}

	result := ProcessData(res)

	expected := "processed data"

	if result != expected {
		t.Errorf("ProcessData() = %s; expected %s\n", result, expected)
	}
}
