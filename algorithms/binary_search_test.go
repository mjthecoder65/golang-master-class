package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	require.Equal(t, 2, BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 52, 80}, 3))
	require.Equal(t, -1, BinarySearch([]int{}, 2))
	require.Equal(t, -1, BinarySearch([]int{30, 40, 52, 60, 71, 90, 150}, 8))
	require.Equal(t, 0, BinarySearch([]int{1}, 1))
	require.Equal(t, 0, BinarySearch([]int{1, 10, 25, 35, 47}, 1))
	require.Equal(t, 8, BinarySearch([]int{10, 22, 33, 49, 51, 62, 75, 82, 97}, 97))
}

func TestBinarySearchRecursive(t *testing.T) {
	require.Equal(t, 2, BinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 52, 80}, 3, 0, 8))
	require.Equal(t, -1, BinarySearchRecursive([]int{}, 2, 0, 0))
	require.Equal(t, -1, BinarySearchRecursive([]int{30, 40, 52, 60, 71, 90, 150}, 8, 0, 6))
	require.Equal(t, 0, BinarySearchRecursive([]int{1}, 1, 0, 0))
	require.Equal(t, 0, BinarySearchRecursive([]int{1, 10, 25, 35, 47}, 1, 0, 4))
	require.Equal(t, 8, BinarySearchRecursive([]int{10, 22, 33, 49, 51, 62, 75, 82, 97}, 97, 0, 8))
}
