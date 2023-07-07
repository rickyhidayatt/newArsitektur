package main

import (
	"fmt"
	netHttp "net/http"
	"os"

	"bni.co.id/xpora/medias/cmd/http"
	"bni.co.id/xpora/medias/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	initLogger()
	runHTTP()
}

func initLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func runHTTP() {
	port := config.GetEnv(config.HTTP_PORT)

	if len(port) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", port))
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Mount("/api", http.CompileRoute(router))
	router.Mount("/debug", middleware.Profiler())

	server := &netHttp.Server{
		Addr:    port,
		Handler: router,
	}
	log.Info(fmt.Sprintf("running bni.co.id/xpora/medias server at http://localhost%s", port))

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
