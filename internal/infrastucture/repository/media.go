package repository

import (
	"time"
)

type Media struct {
	ID        string    `json:"id" db:"id"`
	FileUrl   string    `json:"url" db:"file_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type MediaRepository interface {
	InsertFileMedia(arg *Media) error
	CreateFileMedia(fileBytes []byte, filename, fileURI string) error
}
