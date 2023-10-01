package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinearSearch(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 52, 15, 12, 14, 16, 18, 19, 20, 21, 22, 23, 24}
	require.Equal(t, 2, LinearSearch(elements, 3))
	require.Equal(t, -1, LinearSearch(elements, 8))
	require.Equal(t, 0, LinearSearch(elements, 1))
}

func TestLinearSearchRecursive(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 52, 15, 12, 14, 16, 18, 19, 20, 21, 22, 23, 24}
	require.Equal(t, 2, LinearSearchRecursive(elements, 3, 0))
	require.Equal(t, -1, LinearSearchRecursive(elements, 8, 0))
	require.Equal(t, 0, LinearSearchRecursive(elements, 1, 0))
}
