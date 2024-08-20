package leetcode_50

// URL: https://leetcode.com/problems/valid-anagram/

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
		counter[t[i]]--
	}

	for _, value := range counter {
		if value != 0 {
			return false
		}
	}
	return true
}
