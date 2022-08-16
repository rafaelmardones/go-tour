package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	_, inMap := templateCache[templateName]
	if !inMap {
		fmt.Println("Creating template cache...")
		err := createTemplateCache(templateName)
		if err != nil {
			log.Println(err)
		}
	} else {
		fmt.Println("Retrieving template previously cached...")
	}
	parsedTemplate := templateCache[templateName]
	err2 := parsedTemplate.Execute(w, nil)
	if err2 != nil {
		fmt.Println("error executing template:", err2)
	}
}

func createTemplateCache(t string) error {
	files := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	parsedTemplate, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}
	templateCache[t] = parsedTemplate
	return nil
}
