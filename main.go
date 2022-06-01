package main

import (
	"log"

	"github.com/jason-horvath/goexamples/config"
	"github.com/jason-horvath/goexamples/render"
	"github.com/jason-horvath/goexamples/server"
)

func main() {
	var app config.AppConfig
	tmplCache, err := render.BuildTemplateCache()
	if err != nil {
		log.Println("Unable to load the template cache", err)
	}
	app.TemplateCache = tmplCache
	server.ServerStart()
}
