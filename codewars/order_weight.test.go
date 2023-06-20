package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderWeight(t *testing.T) {
	require.Equal(t, OrderWeight("103 123 4444 99 2000"), "2000 103 123 4444 99")
	require.Equal(t, OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"), "11 11 2000 10003 22 123 1234000 44444444 9999")
	require.Equal(t, OrderWeight(""), "")
}
