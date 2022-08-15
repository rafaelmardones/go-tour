package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, err1 := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.tmpl")
	if err1 != nil {
		fmt.Println("error parsing template:", err1)
		return
	}
	err2 := parsedTemplate.Execute(w, nil)
	if err2 != nil {
		fmt.Println("error executing template:", err2)
	}
}
