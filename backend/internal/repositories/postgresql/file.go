package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Mindslave/skade/backend/internal/entities"
)


//StoreFile is used to store files uploaded to skate
func (r* Repo) StoreFile(ctx context.Context, arg entities.DbStoreFileParams) (error) {
	var file entities.File
	uuid, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error generating uuid: %w", err)
	}
	err = r.Get(&file, "INSERT INTO files VALUES($1, $2, $3, $4, $5) RETURNING *", 
				uuid,
				arg.Filename,
				arg.Filesize,
				arg.FileExtension,
				time.Now())
	if err != nil {
		return fmt.Errorf("error storing file: %w", err)
	}
	return nil
}
