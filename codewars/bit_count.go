package codewars

import "math/bits"

// Question: https://www.codewars.com/kata/526571aae218b8ee490006f4/train/go

func CountBits(number uint) int {

	var counter int

	for number > 0 {
		remainder := number % 2
		if remainder == 1 {
			counter++
		}
		number = number / 2
	}

	return counter
}

// Using built-in function.
func CountBits2(number uint) int {
	return bits.OnesCount(number)
}
