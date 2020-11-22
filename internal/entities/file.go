package entities

import "github.com/google/uuid"

type DbStoreFileParams struct {
	ID			uuid.UUID
	Filename 	string
	Filesize 	int32
}

type File struct {
	ID            uuid.UUID          `json:"id"`
	Filename      string         `json:"filename"`
	Filesize      int32          `json:"filesize"`
}
