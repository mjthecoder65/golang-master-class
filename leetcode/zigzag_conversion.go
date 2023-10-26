package leetcode

// URL: https://leetcode.com/problems/zigzag-conversion/description/
// Title: ZigZag Conversion
// Difficulty: Medium

// Convert : Convert a string to a zigzag pattern.

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := ""

	for row := 0; row < numRows; row++ {
		steps := (numRows - 1) * 2
		for index := row; index < len(s); index += steps {
			result += string(s[index])
			nextCharIndex := index + steps - (row * 2)
			if row > 0 && row < numRows-1 && nextCharIndex < len(s) {
				result += string(s[nextCharIndex])
			}
		}
	}

	return result
}
