package leetcode_50

// URL: https://leetcode.com/problems/intersection-of-two-linked-lists/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)
	currentNode := headA

	for currentNode != nil {
		seen[currentNode] = true
		currentNode = currentNode.Next
	}

	currentNode = headB

	for currentNode != nil {
		if _, ok := seen[currentNode]; ok {
			return currentNode
		}
		currentNode = currentNode.Next
	}

	return nil
}
