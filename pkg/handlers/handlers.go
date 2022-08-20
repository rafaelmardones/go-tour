package handlers

import (
	"net/http"

	"github.com/rafaelmardones/go-tour/pkg/config"
	"github.com/rafaelmardones/go-tour/pkg/models"
	"github.com/rafaelmardones/go-tour/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) {
	Repo = &Repository{
		App: a,
	}
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	var stringMap = make(map[string]string)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	var stringMap = make(map[string]string)
	stringMap["test"] = "Hoi! Hello again!"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
