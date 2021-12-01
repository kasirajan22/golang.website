package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kasirajan22/template/pkg/config"
	"github.com/kasirajan22/template/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(_app *config.AppConfig) {
	app = _app
}

func AddDefaultData(data *models.TemplateData) *models.TemplateData {
	return data
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template

	// get template from cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("cant get error from template cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data)

	_ = t.Execute(buf, data)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("current page is: ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
