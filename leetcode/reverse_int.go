package leetcode

// URL: https://leetcode.com/problems/reverse-integer/description/
// Title: Reverse Integer

const MAX_32_BIT_INT = 2147483647
const MIN_32_BIT_INT = -2147483648

func ReverseInt(number int) int {
	var sign int = 1
	if number < 0 {
		sign = -1
		number *= sign
	}

	reversedNumber := 0

	for number > 0 {
		digit := number % 10
		if reversedNumber*10+digit > MAX_32_BIT_INT || reversedNumber*10+digit < MIN_32_BIT_INT {
			return 0
		}
		reversedNumber = reversedNumber*10 + digit
		number = number / 10
	}

	return sign * reversedNumber
}
