package entities

type File struct {
	ID            int64          `json:"id"`
	Filename      string         `json:"filename"`
	Filesize      int32          `json:"filesize"`
}