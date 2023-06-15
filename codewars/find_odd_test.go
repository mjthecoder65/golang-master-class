package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindOdd(t *testing.T) {
	arr := [...][]int{
		{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
		{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
		{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
		{10},
		{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
		{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
	}

	sol := [...]int{5, -1, 5, 10, 10, 1}

	for i, v := range arr {
		require.Equal(t, sol[i], FindOdd(v))
		require.Equal(t, sol[i], FindOdd2(v))
	}
}
