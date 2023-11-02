package handlers

import (
	"github.com/andriesniemandt/go-webserver/pkg/config"
	"github.com/andriesniemandt/go-webserver/pkg/models"
	"github.com/andriesniemandt/go-webserver/pkg/render"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (c *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.html", &models.TemplateData{})
}

func (c *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	render.Template(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
