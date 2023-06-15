package codewars

// Questionhttps://www.codewars.com/kata/5667e8f4e3f572a8f2000039/train/go

import "strings"

func Accum(str string) string {
	// your code
	var result string

	for index, char := range str {
		firstLetter := strings.ToUpper(string(char))
		restOfLetters := strings.Repeat(strings.ToLower(string(char)), index) // strings.Repeat(str, count int) string

		result += firstLetter + restOfLetters + "-"
	}

	return strings.TrimSuffix(result, "-") // strings.TrimSuffix(str, suffix string) string
}
