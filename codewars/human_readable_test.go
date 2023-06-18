package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHumanReadable(t *testing.T) {
	require.Equal(t, "00:00:00", HumanReadableTime(0))
	require.Equal(t, "00:00:05", HumanReadableTime(5))
	require.Equal(t, "00:01:00", HumanReadableTime(60))
	require.Equal(t, "23:59:59", HumanReadableTime(86399))
	require.Equal(t, "99:59:59", HumanReadableTime(359999))
}
