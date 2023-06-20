package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitString(t *testing.T) {
	require.Equal(t, SplitString("Hello World"), []string{"Hello", "World"})
	require.Equal(t, SplitString("  2000	10003 1234000   44444444   9999 11 11 22 123  "), []string{"2000", "10003", "1234000", "44444444", "9999", "11", "11", "22", "123"})

}

func TestContains(t *testing.T) {
	require.True(t, strings.Contains("Hello", "ll"))
	require.False(t, strings.Contains("Hello", "lll"))
	require.True(t, strings.Contains("Hello", "He"))
}

func TestJoin(t *testing.T) {
	require.Equal(t, strings.Join([]string{"Hello", "World"}, " "), "Hello World")
	require.Equal(t, strings.Join([]string{"Hello", "World"}, ""), "HelloWorld")
	require.Equal(t, strings.Join([]string{"Hello", "World"}, "-"), "Hello-World")
	require.Equal(t, strings.Join([]string{}, " "), "")
}
