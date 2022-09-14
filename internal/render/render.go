package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/rafaelmardones/go-tour/internal/config"
	"github.com/rafaelmardones/go-tour/internal/models"
)

var app *config.AppConfig
var pathToTemplates = "./templates"

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) {
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.CSRFToken = nosurf.Token(r)
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the requested template from cache
	t, ok := tc[templateName]
	if !ok {
		return errors.New("couldn't get template from cache")
	}

	buf := new(bytes.Buffer)

	AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
		return err
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
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
