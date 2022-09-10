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
