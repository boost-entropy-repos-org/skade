package db

import (
	"context"
	"testing"

	"github.com/Mindslave/skade/pkg/generate"
	"github.com/stretchr/testify/require"
)

func createRandomTestUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: generate.RandomUsername(),
		Email: generate.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomTestUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomTestUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.Username, user1.Username)
}




