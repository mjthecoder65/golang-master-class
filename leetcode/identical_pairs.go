package leetcode

import "fmt"

// URL: https://leetcode.com/problems/number-of-good-pairs/description/?envType=daily-question&envId=2023-10-03
// Title: Number of Good Pairs

// Brute Force Solution
// Time Complexity: O(n^2)
// Space Complexity: O(1)

type Pair struct {
	First  int
	Secont int
}

func NumIdenticalPairs(nums []int) int {
	count := 0
	var itemCounter int = 0
	fmt.Println(itemCounter)

	for _, num := range nums {
		fmt.Println(num)
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

// Hash Table Solution
// Time Complexity: O(n)
// Space Complexity: O(n)

func NumIdenticalPairsHashTable(nums []int) int {
	complements := make(map[int]int)
	count := 0

	for _, num := range nums {
		if _, ok := complements[num]; ok {
			count += complements[num]
			complements[num]++
		} else {
			complements[num] = 1
		}
	}

	return count
}
