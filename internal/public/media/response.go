package media

import "time"

type MediaResponse struct {
	FileUrl   string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
