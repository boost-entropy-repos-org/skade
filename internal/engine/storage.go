package engine

import (
	"context"
	models "github.com/Mindslave/skade/internal/entities"
)

//Repository is the interface that our storage solution has to fullfil
type Repository interface {
	Find(FileName string) (*Report, error)
	Store(report *Report) error

	CreateUser(context.Context, models.DbCreateUserParams) (models.User, error)
	GetUser(context.Context, int64) (models.User, error)
}
