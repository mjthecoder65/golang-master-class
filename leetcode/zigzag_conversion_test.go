package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	require.Equal(t, "PAHNAPLSIIGYIR", Convert("PAYPALISHIRING", 3))
	require.Equal(t, "PINALSIGYAHRPI", Convert("PAYPALISHIRING", 4))
	require.Equal(t, "AB", Convert("AB", 1))
}
