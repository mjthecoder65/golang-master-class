package codewars

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToWeird(t *testing.T) {
	require.Equal(t, "ThIs Is A TeSt", toWeirdCase("This is a test"))
	require.Equal(t, "ThIs Is A TeSt", toWeirdCase("ThIs Is A TeSt"))
	require.Equal(t, "WeIrD StRiNg CaSe", toWeirdCase("Weird string case"))
	require.Equal(t, "WeIrD StRiNg CaSe", toWeirdCase("Weird string case"))
}
