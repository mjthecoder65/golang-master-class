package problems

import "strconv"

// URL: https://leetcode.com/problems/number-of-senior-citizens/

func CountSeniors(details []string) int {
	const START, END int = 11, 13
	count := 0

	for _, detail := range details {
		age, _ := strconv.Atoi(detail[START:END])
		if age > 60 {
			count++
		}
	}

	return count
}
