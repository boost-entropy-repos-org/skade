package credentials

import (
	"testing"

	"github.com/Mindslave/skade/backend/pkg/generate"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t* testing.T) {
	password := generate.RandomString(8)
	wrongPassword := generate.RandomString(8)

	hash, err := CreateHash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	
	err = CheckPassword(hash, password)
	require.NoError(t, err)

	err = CheckPassword(hash, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}