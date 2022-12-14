package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rafaelmardones/go-tour/internal/config"
	"github.com/rafaelmardones/go-tour/internal/handlers"
	"github.com/rafaelmardones/go-tour/internal/models"
	"github.com/rafaelmardones/go-tour/internal/render"
)

const port = ":8083"

var app config.AppConfig

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting server on port %s\n", port)

	server := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("error on ListenAndServe:", err)
	}
}

func run() error {
	gob.Register(models.User{})

	// change when in production
	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)
	handlers.NewRepo(&app)
	return nil
}
