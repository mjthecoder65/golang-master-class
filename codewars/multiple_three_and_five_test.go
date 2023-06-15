package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultipleFiveAndThree(t *testing.T) {
	require.Equal(t, 23, Multiple3And5(10))
	require.Equal(t, 78, Multiple3And5(20))
	require.Equal(t, 9168, Multiple3And5(200))
	require.Equal(t, 0, Multiple3And5(-10))
	require.Equal(t, 0, Multiple3And5(0))
}
