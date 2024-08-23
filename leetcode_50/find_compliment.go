package leetcode_50

import (
	"fmt"
	"strconv"
)

func FindComplement(num int) int {
	binary := fmt.Sprintf("%b", num)

	result := ""

	for _, bit := range binary {
		if bit == '1' {
			result += "0"
		} else {
			result += "1"
		}
	}

	number, _ := strconv.ParseInt(result, 2, 64)

	return int(number)
}
