package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindomicNumber(t *testing.T) {
	require.Equal(t, true, IsPalindrome(121))
	require.Equal(t, false, IsPalindrome(-593))
	require.Equal(t, false, IsPalindrome(10))
	require.Equal(t, true, IsPalindrome(0))
}
