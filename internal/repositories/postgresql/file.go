package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Mindslave/skade/internal/entities"
)


//StoreFile is used to store files uploaded to skate
func (r* Repo) StoreFile(ctx context.Context, arg entities.DbStoreFileParams) error {

}