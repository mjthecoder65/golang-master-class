package leetcode_50

import "sort"

// URL: https://leetcode.com/problems/3sum/

func ThreeSum(numbers []int) [][]int {
	result := [][]int{}
	sort.Ints(numbers)

	for i := 0; i < len(numbers); i++ {
		if i > 0 && numbers[i] == numbers[i-1] {
			continue
		}

		left, right := i+1, len(numbers)-1

		for left < right {
			currentSum := numbers[i] + numbers[left] + numbers[right]
			if currentSum == 0 {
				result = append(result, []int{numbers[i], numbers[left], numbers[right]})
				left++
				right--
				// Move the left pointer to next unique number.
				for left < right && numbers[left] == numbers[left-1] {
					left++
				}
				// Move right pointer to next unique numbers
				for left < right && numbers[right] == numbers[right+1] {
					right--
				}
			} else if currentSum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
