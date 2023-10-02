package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWinnerGame(t *testing.T) {
	require.Equal(t, true, WinnerOfGame("AAABABB"))
	require.Equal(t, false, WinnerOfGame("AA"))
	require.Equal(t, false, WinnerOfGame("ABBBBBBBAAA"))
	require.Equal(t, true, WinnerOfGame("AAABB"))
}
