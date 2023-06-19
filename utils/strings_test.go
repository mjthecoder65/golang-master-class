package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitString(t *testing.T) {
	require.Equal(t, SplitString("Hello World"), []string{"Hello", "World"})
	require.Equal(t, SplitString("  2000	10003 1234000   44444444   9999 11 11 22 123  "), []string{"2000", "10003", "1234000", "44444444", "9999", "11", "11", "22", "123"})

}
