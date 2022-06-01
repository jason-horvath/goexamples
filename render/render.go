package render

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/jason-horvath/goexamples/config"
	"github.com/jason-horvath/goexamples/schema"
)

var tmplFunctions = template.FuncMap{}

var app *config.AppConfig

func InitConfig(a *config.AppConfig) {
	app = a
}

// HtmlTemplate - Renders the template using the http response writer
func HtmlTemplate(w http.ResponseWriter, templatePath string, templateData schema.ExampleHtmlData) {
	w.Header().Set("Content-Type", "text/html")
	htmlTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(err)
	}

	err = htmlTemplate.Execute(w, templateData)

	if err != nil {
		panic(err)
	}
}

func CachedHtmlTemplate(w http.ResponseWriter, tmplName string, templateData schema.ExampleHtmlData) {
	tmplCache := app.TemplateCache

	tmpl, ok := tmplCache[tmplName]
	if !ok {
		log.Fatal("Unable to load template from cache.")
	}

	buffer := new(bytes.Buffer)
	_ = tmpl.Execute(w, templateData)
	_, err := buffer.WriteTo(w)
	if err != nil {
		log.Println("Error writing template.", err)
	}
}

func BuildTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.go.html")
	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)

		tmplSet, err := template.New(pageName).Funcs(tmplFunctions).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		layouts, err := filepath.Glob("./templates/layouts/*.go.html")
		log.Println("Layout test", layouts)
		if err != nil {
			return tmplCache, err
		}

		if len(layouts) > 0 {
			tmplSet, err = tmplSet.ParseGlob("./templates/layouts/*.go.html")
			if err != nil {
				return tmplCache, err
			}
		}

		tmplCache[pageName] = tmplSet
	}

	return tmplCache, nil
}

// JsonOutput - Takes the AlteredJson struct and renders it
func JsonOutput(w http.ResponseWriter, jsonToRender schema.OutputJason) {
	w.Header().Set("Content-Type", "application/json")

	renderedJson, err := json.MarshalIndent(jsonToRender, "", "    ")
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(renderedJson)

	if err != nil {
		panic(err)
	}
}
