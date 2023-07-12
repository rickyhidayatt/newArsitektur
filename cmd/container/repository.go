package container

import (
	"bni.co.id/xpora/medias/internal/infrastucture/repository"
	"bni.co.id/xpora/medias/internal/infrastucture/repository/postgres"
)

type RepositoryIoC struct {
	media repository.MediaRepository
}

func NewRepositoryIoC() RepositoryIoC {
	media := postgres.NewMediaPostgres()

	return RepositoryIoC{
		media: media,
	}
}

func (ioc RepositoryIoC) Media() repository.MediaRepository {
	return ioc.media
}
