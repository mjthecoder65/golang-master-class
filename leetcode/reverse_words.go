package leetcode

// October 1, 2023
// URL : https://leetcode.com/problems/reverse-words-in-a-string-iii/

// func reverseWord(word string) string {
// 	var result strings.Builder

// 	for index := len(word) - 1; index >= 0; index-- {
// 		result.WriteByte(word[index])
// 	}
// 	return result.String()
// }

// func reverseWord(word string) string {
// 	runes := []rune(word)
// 	left, right := 0, len(runes)-1

// 	for left < right {
// 		runes[left], runes[right] = runes[right], runes[left]
// 		left++
// 		right--
// 	}

// 	return string(runes)
// }

// func ReverseWords(s string) string {
// 	words := strings.Split(s, " ") // O(n)

// 	for index, word := range words {
// 		words[index] = reverseWord(word) // o(1)
// 	}

// 	return strings.Join(words, " ") // O(n)
// }

func ReverseWords(s string) string {
	runes := []rune(s)
	left := 0
	for right := 0; right < len(runes); right++ {
		if runes[right] == ' ' || right == len(runes)-1 {
			temp_left, temp_right := left, right-1

			if right == len(runes)-1 {
				temp_right = right
			}

			for temp_left < temp_right {
				runes[temp_left], runes[temp_right] = runes[temp_right], runes[temp_left]
				temp_left++
				temp_right--
			}

			left = right + 1
		}
	}
	return string(runes)
}
