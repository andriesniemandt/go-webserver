package handlers

import (
	"github.com/andriesniemandt/go-webserver/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.html")
}
