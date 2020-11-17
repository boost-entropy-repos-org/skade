package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Mindslave/skade/internal/engine"
)

//Repo provides all functions to execute database queries and transactions
type Repo struct {
	*Queries
	db *sql.DB
}

//NewRepo ceates a new Repo
func NewRepo(db *sql.DB) *Repo {
	return &Repo {
		db:		db,
		Queries: New(db),
	}
}

//execTX executes an entire transaction in the database
func (repo *Repo) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil  {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

//Find finds stuff
func (repo *Repo) Find(sth string) (*engine.Report, error) {
	return nil, nil
}

//Store stores stuff
func (repo *Repo) Store(*engine.Report) error {
	return nil
}