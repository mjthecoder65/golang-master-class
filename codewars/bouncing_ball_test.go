package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBouncingBall(t *testing.T) {
	require.Equal(t, 3, BouncingBall(3, 0.66, 1.5))
	require.Equal(t, 15, BouncingBall(30, 0.66, 1.5))
	require.Equal(t, -1, BouncingBall(40, 1, 10))
	require.Equal(t, -1, BouncingBall(10, 0.6, 10))
	require.Equal(t, -1, BouncingBall(40, 1, 40))
}
