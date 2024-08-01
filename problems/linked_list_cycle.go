package problems

// URL: https://leetcode.com/problems/linked-list-cycle/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(N)
// Space Complexity: O(N)
func HasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	currentNode := head

	for currentNode != nil {
		if _, ok := seen[currentNode]; ok {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

// Floyd-Warshal algorithm
// Time complexity: O(N)
// Space complexity: O(1)
func HasCycleOptimal(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
