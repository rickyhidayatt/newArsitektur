package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	kitHttp "github.com/go-kit/kit/transport/http"

	libError "bni.co.id/xpora/medias/internal/server/liberror"
	router "bni.co.id/xpora/medias/internal/transport"
)

func CompileRoute(r *chi.Mux) http.Handler {
	opts := []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(libError.EncodeError),
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(6))

	r.Use(middleware.Timeout(30 * time.Second))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Access-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	router.CompileRoute(r, opts)

	return r
}
