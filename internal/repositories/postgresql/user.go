package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Mindslave/skade/internal/entities"
)

//CreateUser is the Endpoint used to create new users
func (r *Repo) CreateUser(ctx context.Context, arg entities.DbCreateUserParams) (entities.User, error) {
	var u entities.User
	uuid, err := uuid.NewUUID()
	if err != nil {
		return entities.User{}, fmt.Errorf("error generating uuid: %w", err)
	}
	err = r.Get(&u, "INSERT INTO users VALUES($1, $2, $3, $4, $5) RETURNING *", 
				uuid,
				arg.Username,
				arg.hashed_password
				arg.Email,
				time.Now())
	if err != nil {
		return entities.User{}, fmt.Errorf("error creating user: %w", err)
	}
	return u, nil
}

//GetUser is the Endpoint used to get the details of a user
func (r *Repo) GetUser(_ context.Context, id int64) (entities.User, error) {
	var u entities.User
	err := r.Get(&u, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return entities.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}