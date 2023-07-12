package uploader

import "strings"

type MimeType string

const (
	JPEG MimeType = "image/jpeg"
	PNG  MimeType = "image/png"
	PDF  MimeType = "application/pdf"
	MP4  MimeType = "video/mp4"
)

func (m MimeType) GetExtension() string {
	return strings.Split(string(m), "/")[1]
}
