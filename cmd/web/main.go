package main

import (
	"embed"
	"github.com/andriesniemandt/go-webserver/pkg/config"
	"github.com/andriesniemandt/go-webserver/pkg/handlers"
	"github.com/andriesniemandt/go-webserver/pkg/render"
	"log"
	"net/http"
)

const PORT = ":8080"

//go:embed dist
var dist embed.FS

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.Config(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.Handle("/dist/", http.FileServer(http.FS(dist)))

	_ = http.ListenAndServe(PORT, nil)
}
