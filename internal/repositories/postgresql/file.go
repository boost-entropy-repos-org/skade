package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Mindslave/skade/internal/entities"
)


//StoreFile is used to store files uploaded to skate
func (r* Repo) StoreFile(ctx context.Context, arg entities.DbStoreFileParams) (entities.File, error) {
	var file entities.File
	uuid, err := uuid.NewUUID()
	if err != nil {
		return entities.File{}, fmt.Errorf("error generating uuid: %w", err)
	}
	err = r.Get(&u, "INSERT INTO files VALUES($1, $2, $3, $4) RETURNING *", 
				uuid,
				arg.Filename,
				arg.Filesize,
				time.Now())
	if err != nil {
		return entities.User{}, fmt.Errorf("error creating user: %w", err)
	}
	return file, nil
}
}