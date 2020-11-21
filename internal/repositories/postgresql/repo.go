package postgresql

import (
	"github.com/jmoiron/sqlx"
)

//Repo is the storage object
type Repo struct {
	*sqlx.DB
}

//NewRepo creates a new Repo
func NewRepo(db *sqlx.DB) *Repo {
	return &Repo {
		db,
	}
}