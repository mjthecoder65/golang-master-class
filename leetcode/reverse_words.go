package leetcode

import "strings"

// October 1, 2023
// URL : https://leetcode.com/problems/reverse-words-in-a-string-iii/

// func reverseWord(word string) string {
// 	var result strings.Builder

// 	for index := len(word) - 1; index >= 0; index-- {
// 		result.WriteByte(word[index])
// 	}
// 	return result.String()
// }

func reverseWord(word string) string {
	runes := []rune(word)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func ReverseWords(s string) string {
	words := strings.Split(s, " ") // O(n)

	for index, word := range words {
		words[index] = reverseWord(word) // o(1)
	}

	return strings.Join(words, " ") // O(n)
}
