package render

import (
	"net/http"
	"testing"

	"github.com/rafaelmardones/go-tour/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	td := models.TemplateData{}
	r, err := getSession()
	if err != nil {
		t.Fatal(err)
	}
	session.Put(r.Context(), "flash", "testFlash")
	AddDefaultData(&td, r)
	if td.Flash != "testFlash" {
		t.Error("Flash variable could not be set")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/test-url", nil)
	if err != nil {
		return nil, err
	}

	context := r.Context()
	context, err = session.Load(context, r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(context)

	return r, nil
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error rendering home template")
	}

	err = RenderTemplate(ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("didn't return expected error when rendering non existent template")
	}

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
