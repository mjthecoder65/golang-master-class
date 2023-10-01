package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test is reverseVowels

func TestReverseVowel(t *testing.T) {
	require.Equal(t, "holle", ReverseVowels("hello"))
	require.Equal(t, "leotcede", ReverseVowels("leetcode"))
	require.Equal(t, "Aa", ReverseVowels("aA"))
	require.Equal(t, "", ReverseVowels(""))
}
