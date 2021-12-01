package handlers

import (
	"net/http"

	"github.com/kasirajan22/template/pkg/config"
	"github.com/kasirajan22/template/pkg/models"
	"github.com/kasirajan22/template/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//creates a new repository
func NewRepo(_app *config.AppConfig) *Repository {
	return &Repository{
		App: _app,
	}
}

//sets repository for handlers
func NewHandlers(rep *Repository) {
	Repo = rep
}

//func with reciver
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, world"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
