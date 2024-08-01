package leetcode

// URL: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
// TITLE: Two Sum II - Input array is sorted

func twoSumSorted(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
