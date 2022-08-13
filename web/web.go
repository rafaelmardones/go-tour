package main

import (
	"fmt"
	"net/http"
)

const port = ":8083"

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello, this is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello, this is the about page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(":8083", nil)
}
