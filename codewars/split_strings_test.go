package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitString(t *testing.T) {
	require.Equal(t, []string{"ab", "c_"}, SplitString("abc"))
	require.Equal(t, []string{"ab", "cd", "ef"}, SplitString("abcdef"))
	require.Equal(t, []string{"ab", "cd", "ef", "g_"}, SplitString("abcdefg"))
}
