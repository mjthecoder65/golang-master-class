# Golang-master-class
Create this repository for learning  and mastering Golang.
Feel free to contribute.

## Resources
- [Golang Programs](https://www.golangprograms.com/go-language/concurrency.html)
- [Go By Examples](https://gobyexample.com/)


## Problem 1
```go
// Reverse characters of each word in a string.
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

## Problem 2
```go
// Find two numbers whose sum equals the target
func TwoSumHashTable(nums []int, target int) []int {
	complements := make(map[int]int)

	for index, num := range nums {
		if _, ok := complements[num]; ok {
			return []int{complements[num], index}
		} else {
			complements[target-num] = index
		}
	}

	return nil
}
```

## Problem 3
```go
// Reversing Linked List
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var previousNode *ListNode
	currentNode := head

	for currentNode != nil {
		nextTemp := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextTemp
	}

	return previousNode
}
```

## Problem 4
```go
// Merging two sorted Lists
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}
}
```

## Problem 5
```go
func ReverseVowels(s string) string {
    runes := []rune(s)
    left, right := 0, len(s) - 1

    for left < right  {

        // Move the left pointer until it points to vowel
        for left < right && !isVowel(runes[left]) {
            left++
        }

        // Move the right pointer until it points to a vowel

        for left < right && !isVowel(runes[right]) {
            right--
        }
        
        // Swap the vowels
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--

    }

    return string(runes)
}
```