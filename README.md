# Golang-master-class
Create this repository for learning  and mastering Golang.
Feel free to contribute.

## Resources
- [Golang Programs](https://www.golangprograms.com/go-language/concurrency.html)
- [Go By Examples](https://gobyexample.com/)


## Problem 1
```go
func reverseWord(word string) string {
	runes := []rune(word)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func ReverseWords(s string) string {
	words := strings.Split(s, " ")

	for index, word := range words {
		words[index] = reverseWord(word)
	}

	return strings.Join(words, " ")
}
```