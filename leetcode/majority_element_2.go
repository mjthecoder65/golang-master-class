package leetcode

// Day 5
// URL: https://leetcode.com/problems/majority-element-ii/description/?envType=daily-question&envId=2023-10-05

func MajorityElement2(nums []int) []int {
	counter := make(map[int]int)
	elements := []int{}

	for _, num := range nums {
		counter[num]++
	}

	for num, count := range counter {
		if count > len(nums)/3 {
			elements = append(elements, num)
		}
	}

	return elements
}
