package main

import (
	"embed"
	"github.com/andriesniemandt/go-webserver/pkg/config"
	"github.com/andriesniemandt/go-webserver/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

//go:embed dist
var dist embed.FS

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/dist/*", http.FileServer(http.FS(dist)).ServeHTTP)

	return r
}
