package models

import "github.com/rafaelmardones/go-tour/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	CSRFToken string
	Form      *forms.Form
	Data      map[string]interface{}
}
