package codewars

// Question: https://www.codewars.com/kata/55c6126177c9441a570000cc/train/go

import (
	"sort"
	"strings"
)

func SumDigits(number string) int {
	var sum int
	for _, v := range number {
		sum += int(v - '0')
	}
	return sum
}

func OrderWeight(strng string) string {
	words := strings.Fields(strng)

	sort.Slice(words, func(i int, j int) bool {
		if diff := SumDigits(words[i]) - SumDigits(words[j]); diff == 0 {
			return words[i] < words[j]
		} else {
			return diff < 0
		}
	})

	return strings.Join(words, " ")
}
