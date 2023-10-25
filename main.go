package main

import (
	"embed"
	"net/http"
)

const PORT = ":8080"

//go:embed dist
var dist embed.FS

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.Handle("/dist/", http.FileServer(http.FS(dist)))

	_ = http.ListenAndServe(PORT, nil)
}
