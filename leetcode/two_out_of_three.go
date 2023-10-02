package leetcode

// URL: https://leetcode.com/problems/two-out-of-three/
// Title: Two Out of Three

func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {

	counter := make(map[int][]int)

	for _, num1 := range nums1 {
		if _, ok := counter[num1]; !ok {
			counter[num1] = []int{1, 0, 0}
		} else {
			counter[num1][0]++
		}
	}

	for _, num2 := range nums2 {
		if _, ok := counter[num2]; !ok {
			counter[num2] = []int{0, 1, 0}
		} else {
			counter[num2][1]++
		}
	}

	for _, num3 := range nums3 {
		if _, ok := counter[num3]; !ok {
			counter[num3] = []int{0, 0, 1}
		} else {
			counter[num3][2]++
		}
	}

	result := []int{}

	for key, value := range counter {
		if (value[0] > 0 && value[1] > 0) || (value[0] > 0 && value[2] > 0) || (value[1] > 0 && value[2] > 0) {
			result = append(result, key)
		}
	}

	return result
}
