package utils

import "strings"

// Split a string into a slice of strings.
func SplitString(str string) []string {
	return strings.Fields(str)
}

// Join a slice of strings into a single string.
func JoinString(words []string) string {
	return strings.Join(words, " ")
}

// Replace a substring with another substring.
func ReplaceString(str string, old string, new string) string {
	return strings.Replace(str, old, new, -1)
}

// Find index of a substring in a string.
func FindString(str string, substr string) int {
	return strings.Index(str, substr)
}
