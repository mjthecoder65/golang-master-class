package leetcode

// URL: https://leetcode.com/problems/reverse-linked-list/
// Title: Reverse Linked List

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
