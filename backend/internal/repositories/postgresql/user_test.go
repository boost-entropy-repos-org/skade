  
package postgresql

import(
	"testing"
	"context"

	"github.com/Mindslave/skade/backend/internal/entities"
	
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := entities.DbCreateUserParams {
		Username: "TestUser",
		Email: "test@test.com",
		HashedPassword: "abc",
	}

	user, err := testrepo.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
}