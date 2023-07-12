package media

import (
	"bni.co.id/xpora/medias/internal/infrastucture/repository"
	publicMedia "bni.co.id/xpora/medias/internal/public/media"
)

type MediaServiceInterface interface {
	InsertFileMedia(fileUrl string) (*publicMedia.MediaResponse, error)
}

type MediaService struct {
	repository repository.MediaRepository
}

func NewMediaService(mediaRepo repository.MediaRepository) MediaServiceInterface {
	return &MediaService{
		repository: mediaRepo,
	}
}
