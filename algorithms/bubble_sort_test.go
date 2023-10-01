package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 52, 80}, BubbleSort([]int{1, 3, 2, 4, 5, 52, 6, 7, 80}))
	require.Equal(t, []int{}, BubbleSort([]int{}))
	require.Equal(t, []int{2, 5, 6}, BubbleSort([]int{5, 6, 2}))
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 52, 80}, BubbleSort([]int{1, 2, 3, 4, 5, 6, 7, 52, 80}))
	require.Equal(t, []int{5}, BubbleSort([]int{5}))
}
