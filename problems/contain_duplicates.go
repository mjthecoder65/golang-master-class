package problems

// URL: https://leetcode.com/problems/contains-duplicate/

func ContainDuplicates(numbers []int) bool {
	seen := make(map[int]int)
	for index, number := range numbers {
		if _, ok := seen[number]; ok {
			return true
		}
		seen[number] = index
	}
	return false
}
