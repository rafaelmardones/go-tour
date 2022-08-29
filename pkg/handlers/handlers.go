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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	var stringMap = make(map[string]string)
	stringMap["test"] = "Hoi! Hello again!"

	// just to test session storage (retrieval):
	remote_ip := Repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remote_ip

	render.RenderTemplate(w, r, "signup.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) PostSignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form submitted"))
}
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var stringMap = make(map[string]string)
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
