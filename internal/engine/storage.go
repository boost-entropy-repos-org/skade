package engine

import (
	"context"
	"github.com/Mindslave/skade/internal/entities"
)

//Repository is the interface that our storage solution has to fullfil
type Repository interface {
	CreateUser(context.Context, entities.DbCreateUserParams) (entities.User, error)
	GetUser(context.Context, int64) (entities.User, error)
}
