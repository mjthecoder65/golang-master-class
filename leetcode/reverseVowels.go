package leetcode

import "strings"

// URL: https://leetcode.com/problems/reverse-vowels-of-a-string/solutions/
// Title: Reverse Vowels of a String

func isVowel(char rune) bool {
	vowels := "AEIOUaeiou"
	return strings.ContainsRune(vowels, char)
}

func ReverseVowels(s string) string {
	runes := []rune(s)
	left, right := 0, len(s)-1

	for left < right {

		if isVowel(runes[left]) && isVowel(runes[right]) {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}

		if !isVowel(runes[left]) {
			left++
		}

		if !isVowel(runes[right]) {
			right--
		}

	}

	return string(runes)
}

// func ReverseVowels(s string) string {
//     runes := []rune(s)
//     left, right := 0, len(s) - 1

//     for left < right  {

//         // Move the left pointer until it points to vowel
//         for left < right && !isVowel(runes[left]) {
//             left++
//         }

//         // Move the right pointer until it points to a vowel

//         for left < right && !isVowel(runes[right]) {
//             right--
//         }

//         // Swap the vowels
//         runes[left], runes[right] = runes[right], runes[left]
//         left++
//         right--

//     }

//     return string(runes)
// }
