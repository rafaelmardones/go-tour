package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelmardones/go-tour/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error("routes() return is not of type *chi.Mux")
	}
}
