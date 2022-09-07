package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
	"github.com/rafaelmardones/go-tour/internal/config"
	"github.com/rafaelmardones/go-tour/internal/models"
	"github.com/rafaelmardones/go-tour/internal/render"
)

var app config.AppConfig
var pathToTemplates = "./../../templates"

func GetRoutes() http.Handler {
	gob.Register(models.User{})

	// change when in production
	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = true
	render.NewTemplates(&app)
	NewRepo(&app)

	r := chi.NewRouter()

	//r.Use(NoSurf)
	r.Use(middleware.Recoverer)
	r.Use(SessionLoad)
	r.Get("/", Repo.Home)
	r.Get("/signup", Repo.SignUp)
	r.Post("/signup", Repo.PostSignUp)
	r.Post("/check-username-available", Repo.CheckUsernameAvailability)
	r.Get("/login", Repo.Login)
	r.Get("/dashboard", Repo.Dashboard)
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))
	return r

}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	var tc = map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return tc, err
	}
	layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
	if err != nil {
		return tc, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tc, err
		}
		if len(layouts) > 0 {
			t, err = t.ParseFiles(layouts...)
			if err != nil {
				return tc, err
			}
		}
		tc[name] = t
	}
	return tc, nil
}
