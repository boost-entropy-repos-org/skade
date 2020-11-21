package postgresql

import (
	"context"
	"fmt"
	"github.com/Mindslave/skade/internal/entities"
)

func (r *Repo) CreateUser(ctx context.Context, arg entities.DbCreateUserParams) (entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Repo) GetUser(_ context.Context, id int64) (entities.User, error) {
	var u entities.User
	err := r.Get(&u, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return entities.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}