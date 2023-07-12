package endpoint

import (
	"context"
	"database/sql"
	"net/http"

	dbInstance "bni.co.id/xpora/medias/database"
	"bni.co.id/xpora/medias/internal/application"
	mediaVerifications "bni.co.id/xpora/medias/internal/public/media"
	"bni.co.id/xpora/medias/internal/server/database"
	"bni.co.id/xpora/medias/internal/server/liberror"
	"bni.co.id/xpora/medias/internal/server/response"
	"github.com/go-kit/kit/endpoint"
)

func Base64FileUpload(app application.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		payload := req.(*mediaVerifications.Base64UploadRequest)

		err = database.RunInTransaction(ctx, dbInstance.PgDB(), func(ctx context.Context, tx *sql.Tx) error {
			if payload == nil {
				return liberror.New(nil, http.StatusBadRequest, "Invalid request: Payload is nil")
			}

			result, err := app.Commands.UploadBase64.Execute(ctx, *payload)
			if err != nil {
				return liberror.New(err, http.StatusBadRequest, "Invalid request: Failed to execute UploadBase64 command")
			}

			res = result
			return nil
		})

		if err != nil {
			return nil, liberror.New(err, http.StatusBadRequest, "Failed to run transaction")
		}

		return response.ResponseWithRequestTime(ctx, res), nil
	}
}
