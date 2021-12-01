package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kasirajan22/template/pkg/config"
	"github.com/kasirajan22/template/pkg/handlers"
	"github.com/kasirajan22/template/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tmplCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("error creating template")
	}

	app.TemplateCache = tmplCache
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
