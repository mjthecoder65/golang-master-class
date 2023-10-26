package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}
	result := addTwoNumbers(l1, l2)

	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 1 {
		t.Errorf("Expected 7 -> 0 -> 1, got %v", result)
	}

	l1 = &ListNode{Val: 5, Next: nil}
	l2 = &ListNode{Val: 5, Next: nil}
	result = addTwoNumbers(l1, l2)
	if result.Val != 0 || result.Next.Val != 1 {
		t.Errorf("Expected 0 -> 1, got %v", result)
	}
}
