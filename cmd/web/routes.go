package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rafaelmardones/go-tour/internal/config"
	"github.com/rafaelmardones/go-tour/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(NoSurf)
	r.Use(middleware.Recoverer)
	r.Use(SessionLoad)
	r.Get("/", handlers.Repo.Home)
	r.Get("/signup", handlers.Repo.SignUp)
	r.Post("/signup", handlers.Repo.PostSignUp)
	r.Post("/check-username-available", handlers.Repo.CheckUsernameAvailability)
	r.Get("/login", handlers.Repo.Login)
	r.Get("/dashboard", handlers.Repo.Dashboard)
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))
	return r
}
