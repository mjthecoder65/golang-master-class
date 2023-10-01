package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test ReverseWords

func TestReverseWords(t *testing.T) {
	require.Equal(t, "olleh", ReverseWords("hello"))
	require.Equal(t, "", ReverseWords(""))
	require.Equal(t, "s'teL ekat edoCteeL tsetnoc", ReverseWords("Let's take LeetCode contest"))
	require.Equal(t, "doG gniD", ReverseWords("God Ding"))
}
