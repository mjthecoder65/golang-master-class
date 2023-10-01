package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertionSort(t *testing.T) {
	require.Equal(t, []int{1, 2, 5, 8}, InsertionSort([]int{5, 2, 8, 1}))
	require.Equal(t, []int{}, InsertionSort([]int{}))
	require.Equal(t, []int{3, 4, 5}, InsertionSort([]int{4, 3, 5}))
	require.Equal(t, []int{3}, InsertionSort([]int{3}))
}
