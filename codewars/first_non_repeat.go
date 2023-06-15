package codewars

// Question: https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/train/go

import "strings"

func FirstNonRepeating(str string) string {
	counter := make(map[string]int)

	for _, char := range str {
		key := strings.ToLower(string(char))
		counter[key] += 1
	}

	for _, char := range str {
		key := strings.ToLower(string(char))
		if counter[key] == 1 {
			return string(char)
		}
	}

	return ""
}
