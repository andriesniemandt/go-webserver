package main

import (
	"embed"
	"github.com/andriesniemandt/go-webserver/pkg/handlers"
	"net/http"
)

const PORT = ":8080"

//go:embed dist
var dist embed.FS

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.Handle("/dist/", http.FileServer(http.FS(dist)))

	_ = http.ListenAndServe(PORT, nil)
}
