package command

import (
	"context"
	"errors"
	"fmt"

	userDomainService "bni.co.id/xpora/medias/internal/domain/service/media"
	publicMedia "bni.co.id/xpora/medias/internal/public/media"
	"bni.co.id/xpora/medias/internal/server/uploader"
)

type UploadBase64File struct {
	mediaService userDomainService.MediaServiceInterface
}

func NewUploadBase64File(service userDomainService.MediaServiceInterface) UploadBase64File {
	return UploadBase64File{
		mediaService: service,
	}
}

func (c UploadBase64File) Execute(ctx context.Context, payload publicMedia.Base64UploadRequest) (*publicMedia.MediaResponse, error) {

	base64Decode := uploader.Base64File{
		Base64String: payload.Base64String,
		Mime:         uploader.MimeType(payload.Mime),
	}

	// save base64 decode to local storage
	filepath, err := base64Decode.UploadToTempDir()
	if err != nil {
		fmt.Println("Gagal nyimpen ke file dir")
		return nil, errors.New("gagal dir di upload_media")
	}

	// save filepath to db
	saveFile, err := c.mediaService.InsertFileMedia(*filepath)
	if err != nil {
		return nil, errors.New("gagal insert file in upload_media")
	}
	return saveFile, nil
}
