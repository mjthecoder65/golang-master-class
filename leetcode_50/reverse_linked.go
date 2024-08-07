package leetcode_50

// URL: https://leetcode.com/problems/reverse-linked-list/

func ReverseList(head *ListNode) *ListNode {
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
