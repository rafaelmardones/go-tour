package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rafaelmardones/go-tour/pkg/config"
	"github.com/rafaelmardones/go-tour/pkg/handlers"
	"github.com/rafaelmardones/go-tour/pkg/render"
)

const port = ":8083"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(":8083", nil)
}
