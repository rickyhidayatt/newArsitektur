package postgres

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"bni.co.id/xpora/medias/config"
	"bni.co.id/xpora/medias/internal/infrastucture/repository"
	// "bni.co.id/xpora/medias/internal/infrastructure/repository"
)

type MediaStorage struct {
	uploadDir string
}

func NewMediaStorage() repository.MediaRepository {
	return &MediaStorage{
		uploadDir: config.GetEnv(config.TEMP_UPLOAD_DIR),
	}
}

func (p MediaStorage) CreateFileMedia(fileBytes []byte, filename, fileUrl string) error {
	err := os.MkdirAll(p.uploadDir, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(p.uploadDir, filename)
	err = ioutil.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (p MediaStorage) InsertFileMedia(arg *repository.Media) error {
	return errors.New("error unimplemented")
}
