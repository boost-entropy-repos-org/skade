package generate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testRandomString(t *testing.T) {
	string1 := RandomString(8)
	string2 := RandomString(8)
	string3 := RandomString(8)

	require.NotEqual(t, string1, string2)
	require.NotEqual(t, string3, string1)
	require.NotEqual(t, string3, string2)
}