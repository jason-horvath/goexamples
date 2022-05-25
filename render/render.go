package render

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/jason-horvath/goexamples/schema"
)

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
