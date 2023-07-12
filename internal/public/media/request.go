package media

type Base64UploadRequest struct {
	Base64String string `json:"base64_file" validate:"required"`
	Mime         string `json:"mime" validate:"required,mime_type"`
}
