package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Testing IsValidParenthesis

func TestIsValidParenthesis(t *testing.T) {
	require.Equal(t, true, IsValidParenthesis("()"))
	require.Equal(t, true, IsValidParenthesis("()[]{}"))
	require.Equal(t, false, IsValidParenthesis("(]"))
	require.Equal(t, false, IsValidParenthesis("([)]"))
	require.Equal(t, true, IsValidParenthesis("{[]}"))
}
