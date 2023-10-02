package leetcode

// URL: https://leetcode.com/problems/remove-element/description/
// Title: Remove Element

func RemoveElement(nums []int, val int) int {
	remainingElements := []int{}

	for _, num := range nums {
		if num != val {
			remainingElements = append(remainingElements, num)
		}
	}

	// for i := 0; i < len(remainingElements); i++ {
	// 	nums[i] = remainingElements[i]
	// }

	// return len(remainingElements)

	return copy(nums, remainingElements)
}
