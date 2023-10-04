package dsa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashMap(t *testing.T) {
	hashMap := Constructor()
	require.Equal(t, -1, hashMap.Get(1))
	hashMap.Put(1, 10)
	hashMap.Put(2, 34)

	require.Equal(t, 10, hashMap.Get(1))
	require.Equal(t, 34, hashMap.Get(2))

	hashMap.Remove(1)
	require.Equal(t, -1, hashMap.Get(1))
	require.Equal(t, 34, hashMap.Get(2))

}
