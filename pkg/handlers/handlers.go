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
	// just to test session storage:
	remote_ip := r.RemoteAddr
	Repo.App.Session.Put(r.Context(), "remote_ip", remote_ip)

	var stringMap = make(map[string]string)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	var stringMap = make(map[string]string)
	stringMap["test"] = "Hoi! Hello again!"

	// just to test session storage (retrieval):
	remote_ip := Repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remote_ip

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
