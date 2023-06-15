package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitCount(t *testing.T) {
	require.Equal(t, 0, CountBits(0))
	require.Equal(t, 1, CountBits(4))
	require.Equal(t, 5, CountBits(1234))
	require.Equal(t, 2, CountBits(10))
}
