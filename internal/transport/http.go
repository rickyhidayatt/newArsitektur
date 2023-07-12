package transport

import (
	"net/http"

	"bni.co.id/xpora/medias/cmd/container"
	"bni.co.id/xpora/medias/internal/endpoint"
	publicUpload "bni.co.id/xpora/medias/internal/public/media"
	"bni.co.id/xpora/medias/internal/server"
	"github.com/go-chi/chi/v5"
	kitHttp "github.com/go-kit/kit/transport/http"
)

// Root
func CompileRoute(r *chi.Mux, opts []kitHttp.ServerOption) http.Handler {
	r.Post("/upload/media", uploadMedia(opts))
	return r
}

func uploadMedia(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// panggil endpoint
		server.NewHttpServer(endpoint.Base64FileUpload(container.Injector().Application.Application()),
			server.Option{
				DecodeModel: &publicUpload.Base64UploadRequest{},
			},
			opts,
		).ServeHTTP(w, r)
	}
}
