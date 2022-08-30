package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/rafaelmardones/go-tour/internal/config"
	"github.com/rafaelmardones/go-tour/internal/models"
	"github.com/rafaelmardones/go-tour/internal/render"
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

type checkUsernameAvailabilityJson struct {
	Username  string `json:"username"`
	Available bool   `json:"available"`
}

func (m *Repository) CheckUsernameAvailability(w http.ResponseWriter, r *http.Request) {
	username := r.Form.Get("username")
	notAvailableUsernames := []string{"signup", "login", "dashboard"}
	available := true
	for _, forbiddenUsername := range notAvailableUsernames {
		if strings.ToLower(username) == forbiddenUsername {
			available = false
		}
	}
	resp := checkUsernameAvailabilityJson{
		Username:  username,
		Available: available,
	}
	out, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error in CheckUsernameAvailability: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var stringMap = make(map[string]string)
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
