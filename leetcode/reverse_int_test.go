package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	require.Equal(t, 321, ReverseInt(123))
	require.Equal(t, -321, ReverseInt(-123))
	require.Equal(t, 21, ReverseInt(120))
	require.Equal(t, 0, ReverseInt(1534236469))
}
