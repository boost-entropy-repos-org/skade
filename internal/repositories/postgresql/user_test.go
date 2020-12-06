package postgresql

import(
	"testing"
	"context"

	"github.com/Mindslave/skade/internal/entities"
	
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := entities.DbCreateUserParams {
		Username: "TestUser",
		Email: "test@test.com",
		HashedPassword: "abc",
	}

	_, err := testrepo.CreateUser(context.Background(), arg)
	require.NoError(t, err)
}