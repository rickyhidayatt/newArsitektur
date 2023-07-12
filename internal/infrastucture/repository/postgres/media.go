package postgres

import (
	"database/sql"
	"errors"

	"bni.co.id/xpora/medias/database"
	"bni.co.id/xpora/medias/internal/infrastucture/repository"
	// "bni.co.id/xpora/medias/internal/infrastructure/repository"
)

type mediaPostgres struct {
	db *sql.DB
}

func NewMediaPostgres() repository.MediaRepository {
	return &mediaPostgres{
		db: database.PgDB(),
	}
}

func (p mediaPostgres) InsertFileMedia(arg *repository.Media) error {
	stmt, err := p.db.Prepare("INSERT INTO medias (id, file_url, created_at) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(arg.ID, arg.FileUrl, arg.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p mediaPostgres) CreateFileMedia(fileBytes []byte, filename, fileUrl string) error {
	return errors.New("error unimplemented")
}
