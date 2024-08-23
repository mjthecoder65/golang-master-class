package leetcode_50

import "sort"

// URL: https://leetcode.com/problems/group-anagrams/

func sortString(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Custom type
type stringSlice []string

func GroupAnagrams(words []string) [][]string {
	result := [][]string{}
	groups := make(map[string]stringSlice)

	for _, word := range words {
		key := sortString(word)
		groups[key] = append(groups[key], word)
	}

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
