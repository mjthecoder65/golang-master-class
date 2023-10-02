# Golang-master-class
Create this repository for learning  and mastering Golang.
Feel free to contribute.

## Resources
- [Golang Programs](https://www.golangprograms.com/go-language/concurrency.html)
- [Go By Examples](https://gobyexample.com/)


## Problem 1
```go
// Reverse characters of each word in a string.
func ReverseWords(s string) string {
	runes := []rune(s)
	left := 0
	for right := 0; right < len(runes); right++ {
		if runes[right] == ' ' || right == len(runes)-1 {
			temp_left, temp_right := left, right-1

			if right == len(runes)-1 {
				temp_right = right
			}

			for temp_left < temp_right {
				runes[temp_left], runes[temp_right] = runes[temp_right], runes[temp_left]
				temp_left++
				temp_right--
			}

			left = right + 1
		}
	}
	return string(runes)
}
```

## Problem 2
```go
// Find two numbers whose sum equals the target
func TwoSum(nums []int, target int) []int {
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

## Problem 6
```go
func WinnerOfGame(colors string) bool {
	aliceCount, bobCount := 0, 0

	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == colors[i-1] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				aliceCount++
			} else {
				bobCount++
			}
		}

	}
	return aliceCount > bobCount
}
```