package codewars

// Question: https://www.codewars.com/kata/52b757663a95b11b3d00062d/train/go

import "strings"

func convertToWeird(word string) string {
	var result string

	for index, charCode := range word {
		var letter string
		if index%2 == 0 {
			letter = strings.ToUpper(string(charCode))
		} else {
			letter = strings.ToLower(string(charCode))
		}

		result += letter
	}

	return result
}

func toWeirdCase(str string) string {
	var result string

	words := strings.Split(str, " ")

	for _, word := range words {
		weirdString := convertToWeird(word)
		result += weirdString + " "
	}

	return strings.TrimSuffix(result, " ")
}
