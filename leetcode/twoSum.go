package leetcode

// URL: https://leetcode.com/problems/two-sum/submissions/
// Title: Two Sum Problem

// Brute Force Solution
// Time Complexity: O(n^2)
// Space Complexity: O(1)

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Hash Table Solution
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSumHashTable(nums []int, target int) []int {
	complements := make(map[int]int)

	for index, num := range nums {
		if _, ok := complements[num]; ok {
			return []int{complements[num], index}
		} else {
			complements[target-num] = index
		}
	}

	return nil
}
