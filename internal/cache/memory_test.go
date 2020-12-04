package cache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemoryCachePut(t *testing.T) {
	memory := Memory()

	const key = "test"
	const expected = true

	err := memory.Set(key, expected)
	require.NoError(t, err)

	value, present, err := memory.Get(key)
	require.NoError(t, err)
	require.True(t, present)
	require.Equal(t, value, expected)
}
