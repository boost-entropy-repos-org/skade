package engine

import (
	"database/sql"
)

type User struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Report struct {
	FileName  string
	FileSize  float64
	FileType  string
	Malicious bool
}