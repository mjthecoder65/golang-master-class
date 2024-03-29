package codewars

//Question: https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

// Solution1: Using xor operator
func FindOdd(seq []int) int {
	// Using xor operator.
	// 0 ^ 0 = 0
	// 1 ^ 1 = 0
	// 0 ^ 1 = 1
	// 1 ^ 0 = 1

	var result int

	for _, num := range seq {
		result ^= num
	}

	return result
}

// Solution2: Using map
func FindOdd2(seq []int) int {
	counter := make(map[int]int)

	for index, num := range seq {
		counter[num]++
		if index == len(seq)-1 && counter[num]%2 != 0 {
			return num
		}

		if counter[num] > 1 && counter[num]%2 != 0 {
			return num
		}

	}

	return 0
}
