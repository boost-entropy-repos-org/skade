package entities

import (
	"time"

	"github.com/google/uuid"
)

//DbStoreFileParams contains all Parameters required to store a file
type DbStoreFileParams struct {
	ID				uuid.UUID
	Filename 		string
	Filesize 		int64
	FileExtension 	string
}

//File is the struct representing a File uploaded to skade
type File struct {
	ID            uuid.UUID  `json:"id"`
	Filename      string     `json:"filename"`
	Filesize      int64      `json:"filesize"`
	FileExtension string	 `json:"file_extension" db:"fileextension"`
	UploadedAt	  time.Time	 `json:"uploaded_at"`
}
