package leetcode

import "sort"

//  URL : https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/?envType=daily-question&envId=2023-10-10

func findMaxElement(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getUniqueElements(elements []int) []int {
	hashMap := make(map[int]bool)
	uniqueElements := []int{}

	for _, element := range elements {
		if _, ok := hashMap[element]; !ok {
			uniqueElements = append(uniqueElements, element)
			hashMap[element] = true
		}
	}

	return uniqueElements
}

func minOperations(nums []int) int {
	length := len(nums)
	nums = getUniqueElements(nums)
	sort.Ints(nums)

	right, maxOps := 0, 0
	for left := range nums {
		for right < len(nums) && nums[right]-nums[left] < length {
			right++
		}
		maxOps = findMaxElement(maxOps, right-left)
	}

	return length - maxOps
}
