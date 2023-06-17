package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayDiff(t *testing.T) {
	require.Equal(t, []int{2}, ArrayDiff([]int{1, 2}, []int{1}))
	require.Equal(t, []int{2, 2}, ArrayDiff([]int{1, 2, 2}, []int{1}))
	require.Equal(t, []int{1}, ArrayDiff([]int{1, 2, 2}, []int{2}))
}
