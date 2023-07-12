package uploader

import (
	"os"
)

type StorageCloudUploader interface {
	UploadToStorage(folder string) (url *string, err error)
}

type FileUpload struct {
	Mime     MimeType
	FileName string
	File     os.File
}
