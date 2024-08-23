package leetcode_50

// URL: https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func LengthOfLongestSubstring(s string) int {
	hashMap := make(map[byte]int)
	left := 0
	maxLength := 0

	for right := 0; right < len(s); right++ {

		key := s[right]
		index, ok := hashMap[key]

		if !ok || index < left {
			maxLength = Max(maxLength, right-left+1)
		} else {
			left = hashMap[key] + 1
		}
		hashMap[key] = right
	}

	return maxLength
}
