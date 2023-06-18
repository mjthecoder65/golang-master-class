package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	t.Run("should return [0, 2]", func(t *testing.T) {
		got := TwoSum([]int{1, 2, 3}, 4)
		want := [2]int{0, 2}
		require.Equal(t, got, want)
	})

	t.Run("should return [1, 2]", func(t *testing.T) {
		got := TwoSum([]int{1234, 5678, 9012}, 14690)
		want := [2]int{1, 2}
		require.Equal(t, got, want)
	})

	t.Run("should return [0, 1]", func(t *testing.T) {
		got := TwoSum([]int{2, 2, 3}, 4)
		want := [2]int{0, 1}
		require.Equal(t, got, want)
	})

	t.Run("should return []", func(t *testing.T) {
		got := TwoSum([]int{}, 4)
		want := [2]int{}
		require.Equal(t, got, want)
	})
}
