package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/rafaelmardones/go-tour/pkg/config"
	"github.com/rafaelmardones/go-tour/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) {
	td.CSRFToken = nosurf.Token(r)
	td.StringMap["test"] = "Override!"
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the requested template from cache
	t, ok := tc[templateName]
	if !ok {
		log.Fatal("Couldn't get the requested template from cache")
	}

	buf := new(bytes.Buffer)

	AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	log.Println("Creating template cache")
	var tc = map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tc, err
	}
	layouts, err := filepath.Glob("./templates/*.layout.tmpl")
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
