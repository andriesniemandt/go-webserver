package main

import (
	"embed"
	"github.com/andriesniemandt/go-webserver/pkg/config"
	"github.com/andriesniemandt/go-webserver/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//go:embed dist
var dist embed.FS

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/dist/*", http.FileServer(http.FS(dist)).ServeHTTP)

	return mux
}
