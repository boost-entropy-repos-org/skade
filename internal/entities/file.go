package entities

import (
	"time"

	"github.com/google/uuid"
)

type DbStoreFileParams struct {
	ID			uuid.UUID
	Filename 	string
	Filesize 	int64
}

type File struct {
	ID            uuid.UUID  `json:"id"`
	Filename      string     `json:"filename"`
	Filesize      int64      `json:"filesize"`
	UploadedAt	  time.Time	 `json:"uploaded_at"`
}
