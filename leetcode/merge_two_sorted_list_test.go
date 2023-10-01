package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test merge two sorted list

func TestMergeTwoSortedList(t *testing.T) {
	require.Equal(t, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}, MergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
	require.Nil(t, MergeTwoLists(nil, nil))
}
