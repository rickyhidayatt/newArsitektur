package media

import (
	"errors"
	"log"
	"strings"
	"time"

	"bni.co.id/xpora/medias/internal/infrastucture/repository"
	publicMedia "bni.co.id/xpora/medias/internal/public/media"
	"github.com/google/uuid"
)

func (s *MediaService) InsertFileMedia(fileUrl string) (*publicMedia.MediaResponse, error) {
	if fileUrl == "" {
		return nil, errors.New("file url is empety")
	}

	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	media := repository.Media{
		ID:        id,
		FileUrl:   fileUrl,
		CreatedAt: time.Now(),
	}

	err := s.repository.InsertFileMedia(&media)
	if err != nil {
		log.Fatal(err)
	}

	fileResponse := publicMedia.MediaResponse{
		FileUrl:   media.FileUrl,
		CreatedAt: media.CreatedAt,
	}

	return &fileResponse, nil
}
