package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {

	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get the requested template from cache
	t, ok := tc[templateName]
	if !ok {
		log.Fatal("Couldn't get the requested template from cache")
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	var tc = map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tc, err
	}
	layouts, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return tc, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tc, err
		}
		if len(layouts) > 0 {
			t, err = t.ParseFiles(layouts...)
			if err != nil {
				return tc, err
			}
		}
		tc[name] = t
	}
	return tc, nil
}
