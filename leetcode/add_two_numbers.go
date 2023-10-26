package leetcode

// URL: https://leetcode.com/problems/add-two-numbers/description/
// Title: Add Two Numbers
// Difficulty: Medium

// ListNode : Definition for singly-linked list.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, currentNode *ListNode
	var digit, carry int

	for l1 != nil || l2 != nil {
		if l2 == nil && l1 != nil {
			digit = (l1.Val + carry) % 10
			carry = (l1.Val + carry) / 10
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			digit = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2 = l2.Next
		} else {
			digit = (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
			l1, l2 = l1.Next, l2.Next
		}

		node := &ListNode{Val: digit, Next: nil}

		if head == nil {
			head = node
			currentNode = head
		} else {
			currentNode.Next = node
			currentNode = node
		}

	}

	if carry != 0 {
		node := &ListNode{Val: carry, Next: nil}
		currentNode.Next = node
	}

	return head
}
