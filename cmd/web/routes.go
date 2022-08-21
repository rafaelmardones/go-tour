package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rafaelmardones/go-tour/pkg/config"
	"github.com/rafaelmardones/go-tour/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(NoSurf)
	r.Use(middleware.Recoverer)
	r.Use(SessionLoad)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))
	return r
}
