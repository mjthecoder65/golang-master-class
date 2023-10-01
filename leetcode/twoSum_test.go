package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	require.Equal(t, []int{0, 1}, TwoSum([]int{2, 7, 11, 15}, 9))
	require.Equal(t, []int{1, 2}, TwoSum([]int{3, 2, 4}, 6))
	require.Nil(t, TwoSum([]int{3, 5}, 6))
}

func TestTwoSumHashTable(t *testing.T) {
	require.Equal(t, []int{0, 1}, TwoSumHashTable([]int{2, 7, 11, 15}, 9))
	require.Equal(t, []int{1, 2}, TwoSumHashTable([]int{3, 2, 4}, 6))
	require.Nil(t, TwoSumHashTable([]int{3, 5}, 6))
}
