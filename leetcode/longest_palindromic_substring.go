package leetcode

// URL: https://leetcode.com/problems/longest-palindromic-substring/
// Title: Longest Palindromic Substring
// Difficulty: Medium

// Time complexity: O(n)
// Space complexity: O(1)
func isPalindrome(word string) bool {
	left, right := 0, len(word)-1

	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func longestPalindrome(s string) string {
	maxLength := 0
	longestPalindrome := ""

	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			word := s[i:j]

			if isPalindrome(word) && len(word) > maxLength {
				maxLength = len(word)
				longestPalindrome = word
			}
		}
	}

	return longestPalindrome
}

// Time complexity: O(n)
// Space complexity: O(1)
// Using sliding window technique
