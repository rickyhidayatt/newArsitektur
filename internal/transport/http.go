package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	kitHttp "github.com/go-kit/kit/transport/http"
)

func CompileRoute(r *chi.Mux, opts []kitHttp.ServerOption) http.Handler {

	r.Get("/health-check", healthCheck(opts))

	return r
}

func healthCheck(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("it's alive :)"))
	}
}
