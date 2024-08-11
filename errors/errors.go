package errors

import (
	"errors"
	"fmt"
	"math"
)

/*
	Go does not raise error, instead functions return error of type
		type error interface {
			Error() string
		}

	error can be returned as nil or using errors.New() to construct basic error message
	Note: error messages should in lower case, otherwise panic.

	type errors struct {
		Message String
	}

	func (error *errors) New (message string) *errors {
		error.Message = message
		return errors{Message: message}
	}
*/

func Sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("math: numbers must be greater than zero")
	}
	return math.Sqrt(number), nil // Always result first, and then error next.
}

func LearnErrorHandling() {

	result, err := Sqrt(9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
