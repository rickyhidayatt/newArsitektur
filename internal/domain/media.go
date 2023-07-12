package domain

import (
	"time"

	"bni.co.id/xpora/medias/internal/infrastucture/repository"
	publicMedia "bni.co.id/xpora/medias/internal/public/media"
	"bni.co.id/xpora/medias/internal/server/encoding"
)

// keseluruhan yang di pake di response dan juga di pake di db
type Media struct {
	ID        string    `json:"id"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"created_at"`
	// FileBytes []byte    `json:"file_byte"`
}

func (u *Media) FromPublicModel(userPublic interface{}) {
	_ = encoding.TransformObject(userPublic, u)
}

func (u *Media) ToPublicModel() *publicMedia.MediaResponse {
	mediaPublic := &publicMedia.MediaResponse{}
	_ = encoding.TransformObject(u, mediaPublic)
	return mediaPublic
}

func (u *Media) FromRepositoryModel(mediaRepo interface{}) {
	_ = encoding.TransformObject(mediaRepo, u)
}

func (u *Media) ToRepositoryModel() *repository.Media {
	mediaRepo := &repository.Media{}
	_ = encoding.TransformObject(u, mediaRepo)
	return mediaRepo
}
