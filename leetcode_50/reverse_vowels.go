package leetcode_50

func IsVowel(charCode rune) bool {
	return charCode == 'a' || charCode == 'e' || charCode == 'i' || charCode == 'o' || charCode == 'u' ||
		charCode == 'A' || charCode == 'E' || charCode == 'I' || charCode == 'O' || charCode == 'U'
}

func ReverseVowels(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		// Move the left pointer until your find the vowel.
		for left < right && !IsVowel(runes[left]) {
			left++
		}

		// Move the right pointer until you find the vowel.
		for left < right && !IsVowel(runes[right]) {
			right--
		}

		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
