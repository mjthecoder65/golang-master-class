package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstNonRepeating(t *testing.T) {
	require.Equal(t, "a", FirstNonRepeating("a"))
	require.Equal(t, "t", FirstNonRepeating("stress"))
	require.Equal(t, "e", FirstNonRepeating("moonmen"))
	require.Equal(t, ",", FirstNonRepeating("Go hang a salami, I'm a lasagna hog!"))
}
