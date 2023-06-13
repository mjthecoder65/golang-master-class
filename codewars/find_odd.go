package codewars

//Question: https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

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

func FindOdd2(seq []int) int {
	counter := make(map[int]int)

	for _, num := range seq {
		counter[num]++
	}

	for num, count := range counter {
		if count%2 != 0 {
			return num
		}
	}

	return 0
}
