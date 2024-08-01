package problems

// URL: https://leetcode.com/problems/two-sum/

func TwoSum(numbers []int, target int) []int {
	compliments := make(map[int]int)

	for currentIndex, number := range numbers {
		if prevIndex, ok := compliments[number]; ok {
			return []int{prevIndex, currentIndex}
		}
		compliments[target-number] = currentIndex // target- number, the number we are looking in the future.
	}

	return []int{}
}
