package codewars

func TwoSum(numbers []int, target int) [2]int {
	complements := make(map[int]int)

	for i, number := range numbers {
		key := target - number
		if index, ok := complements[key]; ok {
			return [2]int{index, i}
		}
		complements[number] = i
	}

	return [2]int{}

}
