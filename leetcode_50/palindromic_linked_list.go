package leetcode_50

// URL: https://leetcode.com/problems/palindrome-linked-list/

func IsPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	firstHalf := head
	secondHalf := reverse(slow)

	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prevNode *ListNode
	currentNode := head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	return prevNode
}
