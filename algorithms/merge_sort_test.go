package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSort(t *testing.T) {
	require.Equal(t, []int{1, 2, 5, 8}, MergeSort([]int{5, 2, 8, 1}))
	require.Equal(t, []int{1, 2, 5, 8}, MergeSort([]int{1, 2, 5, 8}))
	require.Equal(t, []int{}, MergeSort([]int{}))
	require.Equal(t, []int{3, 4, 5}, MergeSort([]int{4, 3, 5}))
	require.Equal(t, []int{3}, MergeSort([]int{3}))
}
