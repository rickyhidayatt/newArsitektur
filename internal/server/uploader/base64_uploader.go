package uploader

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Base64File struct {
	Base64String string
	Mime         MimeType
	fileBuffer   interface{}
}

func (b *Base64File) base64ToBuffer() error {
	b64Decode, err := base64.StdEncoding.DecodeString(b.Base64String)
	if err != nil {
		return errors.New("invalid base64 type")
	}

	switch b.Mime {
	case PNG:
		{
			var src = new(bytes.Buffer)
			pngI, err := png.Decode(bytes.NewReader([]byte(b64Decode)))
			if err != nil {
				return err
			}
			err = png.Encode(src, pngI)
			if err != nil {
				return err
			}
			b.fileBuffer = src
		}
	case JPEG:
		{
			var src = new(bytes.Buffer)
			jpgI, err := jpeg.Decode(bytes.NewReader([]byte(b64Decode)))
			if err != nil {
				return err
			}
			err = jpeg.Encode(src, jpgI, nil)
			if err != nil {
				return err
			}
			b.fileBuffer = src
		}
	default:
		{
			b.fileBuffer = b64Decode
		}
	}

	return nil
}

// func (b *Base64File) UploadToTempDir() (*string, error) {
// 	if err := b.base64ToBuffer(); err != nil {
// 		return nil, err
// 	}

// 	fileDirectory := config.GetEnv(config.TEMP_UPLOAD_DIR)
// 	fileExt := b.Mime.GetExtension()
// 	fileName := fmt.Sprintf(`%s-%d`, uuid.NewString(), time.Now().Unix())
// 	filePath := fmt.Sprintf("%s%c%s.%s", fileDirectory, os.PathSeparator, fileName, fileExt)

// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// defer file.Close()

// 	switch b.fileBuffer.(type) {
// 	case []byte:
// 		if _, err = file.Write(b.fileBuffer.([]byte)); err != nil {
// 			return nil, err
// 		}
// 	case *bytes.Buffer:
// 		if _, err = io.Copy(file, b.fileBuffer.(*bytes.Buffer)); err != nil {
// 			return nil, err
// 		}
// 	}

// 	return &filePath, nil
// }

func (b *Base64File) UploadToTempDir() (*string, error) {
	if err := b.base64ToBuffer(); err != nil {
		return nil, err
	}

	fileDirectory := "C:\\Users\\User\\Documents\\XporaMedia" // Ganti dengan direktori yang diinginkan
	fileExt := b.Mime.GetExtension()
	fileName := fmt.Sprintf(`%s-%d`, uuid.NewString(), time.Now().Unix())
	filePath := filepath.Join(fileDirectory, fileName+"."+fileExt)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	switch b.fileBuffer.(type) {
	case []byte:
		if _, err = file.Write(b.fileBuffer.([]byte)); err != nil {
			return nil, err
		}
	case *bytes.Buffer:
		if _, err = io.Copy(file, b.fileBuffer.(*bytes.Buffer)); err != nil {
			return nil, err
		}
	}

	return &filePath, nil
}
