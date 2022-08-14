package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8083"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello, this is the about page")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, err1 := template.ParseFiles("./templates/" + templateName)
	if err1 != nil {
		fmt.Println("error parsing template:", err1)
		return
	}
	err2 := parsedTemplate.Execute(w, nil)
	if err2 != nil {
		fmt.Println("error executing template:", err2)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(":8083", nil)
}
