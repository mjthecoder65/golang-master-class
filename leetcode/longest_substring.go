package leetcode

// URL: https://leetcode.com/problems/longest-substring-without-repeating-characters/
// TITLE: Longest Substring Without Repeating Characters

// func isUnique(word string) bool {
// 	hashMap := map[rune]bool{}

// 	for _, char := range word {
// 		if hashMap[char] {
// 			return false
// 		}
// 		hashMap[char] = true
// 	}
// 	return true
// }

// func getSubstrings(word string) []string {
// 	substrings := []string{}

// 	for i := 0; i < len(word); i++ {
// 		for j := i + 1; j <= len(word); j++ {
// 			substrings = append(substrings, word[i:j])
// 		}
// 	}

// 	return substrings
// }

// func lengthOfLongestSubstring(s string) int {
// 	substrings := getSubstrings(s)
// 	longest := 0

// 	for _, substring := range substrings {
// 		if isUnique(substring) && len(substring) > longest {
// 			longest = len(substring)
// 		}
// 	}

// 	return longest
// }

// Using a sliding window approach
func lengthOfLongestSubstring(s string) int {
	hashMap := map[rune]int{}
	longest := 0
	windowStart, windowEnd := 0, 0

	for windowEnd < len(s) {
		char := rune(s[windowEnd])
		if index, ok := hashMap[char]; ok {
			windowStart = max(windowStart, index+1)
		}
		hashMap[char] = windowEnd
		longest = max(longest, windowEnd-windowStart+1)
		windowEnd++
	}

	return longest
}
