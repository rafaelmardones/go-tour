package main

import (
	"fmt"
	"net/http"

	"github.com/rafaelmardones/go-tour/pkg/handlers"
)

const port = ":8083"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(":8083", nil)
}
