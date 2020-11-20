// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type File struct {
	ID            int64          `json:"id"`
	Filename      string         `json:"filename"`
	Filesize      int32          `json:"filesize"`
	Fileextension sql.NullString `json:"fileextension"`
	UploadedAt    sql.NullTime   `json:"uploaded_at"`
	UploadedBy    sql.NullInt64  `json:"uploaded_by"`
}

type Report struct {
	ID        int64         `json:"id"`
	File      sql.NullInt64 `json:"file"`
	Malicious bool          `json:"malicious"`
}

type User struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
}