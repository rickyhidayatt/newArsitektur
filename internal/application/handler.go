package application

import (
	"bni.co.id/xpora/medias/internal/application/command"
	mediaService "bni.co.id/xpora/medias/internal/domain/service/media"
)

type Commands struct {
	UploadBase64 command.UploadBase64File
}

type Queries struct{}

type Application struct {
	Commands Commands
	Queries  Queries
}

func New(
	mediaSv mediaService.MediaServiceInterface,
) Application {
	return Application{
		Commands: Commands{
			UploadBase64: command.NewUploadBase64File(mediaSv),
		},
		Queries: Queries{},
	}
}
