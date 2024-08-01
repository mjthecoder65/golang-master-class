package leetcode

// URL: https://leetcode.com/problems/length-of-last-word/

func LengthOfLastWord(str string) int {
	end := len(str) - 1
	counter := 0
	for end >= 0 {
		if str[end] == ' ' {
			if counter > 0 {
				return counter
			}
		} else {
			counter++
		}
		end--
	}
	return counter
}
