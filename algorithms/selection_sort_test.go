package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	require.Equal(t, []int{1, 2, 5, 8}, SelectionSort([]int{5, 2, 8, 1}))
	require.Equal(t, []int{1, 2, 5, 8}, SelectionSort([]int{1, 2, 5, 8}))
	require.Equal(t, []int{}, SelectionSort([]int{}))
	require.Equal(t, []int{3, 4, 5}, SelectionSort([]int{4, 3, 5}))
	require.Equal(t, []int{3}, SelectionSort([]int{3}))
}
