package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseList(t *testing.T) {
	require.Equal(t, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}, ReverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}))
	require.Equal(t, &ListNode{1, nil}, ReverseList(&ListNode{1, nil}))
	require.Nil(t, ReverseList(nil))
}
