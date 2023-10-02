package leetcode

// URL: https://leetcode.com/problems/palindrome-number/description/
// Title: Palindrome Number

// func IsPalindrome(number int) bool {

// 	if number < 0 {
// 		// Negative number cannot be a palindrome
// 		return false
// 	}

// 	digits := []int{}

// 	for number > 0 {
// 		digit := number % 10
// 		digits = append(digits, digit)
// 		number = number / 10
// 	}

// 	left, right := 0, len(digits)-1

// 	for left < right {
// 		if digits[left] != digits[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}

// 	return true

// }

func IsPalindrome(number int) bool {
	if number < 0 {
		return false
	}

	originalNumber := number
	reversedNumber := 0

	for number > 0 {
		digit := number % 10
		reversedNumber = reversedNumber*10 + digit
		number = number / 10
	}

	return originalNumber == reversedNumber
}
