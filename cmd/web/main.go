package main

import (
	"github.com/andriesniemandt/go-webserver/pkg/config"
	"github.com/andriesniemandt/go-webserver/pkg/handlers"
	"github.com/andriesniemandt/go-webserver/pkg/render"
	"log"
	"net/http"
)

const PORT = ":8080"

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

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
