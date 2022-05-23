package server

import (
	"fmt"
	"net/http"
	"text/template"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type ExampleHtmlData struct {
	Heading     string
	Description string
	ListItems   []string
}

// ServerStart - Start the server after handling the routes.
func ServerStart() {
	http.HandleFunc("/", index)

	uri, handler := multiReturnHandler()
	http.HandleFunc(uri, handler)

	htmlUri, htmlHandler := htmlExample()
	http.HandleFunc(htmlUri, htmlHandler)

	fmt.Println("The server is listening. Go to http://localhost:3500 in your browser")
	_ = http.ListenAndServe(":3500", nil)
}

// index - The handler for the index route
func index(w http.ResponseWriter, r *http.Request) {
	content, err := fmt.Fprintf(w, "This is the index handler in the server package.")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes in the response content: %d\n", content)
}

// multiReturnHandler - An example of using a multiple return function as a handler
func multiReturnHandler() (string, Handler) {
	uri := "/multireturnhandler"
	handler := func(w http.ResponseWriter, r *http.Request) {
		content, err := fmt.Fprintf(w, "This is the multi return handler in the server package.")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Number of bytes in the response content: %d\n", content)
	}

	return uri, handler
}

// htmlExample - Html Template route hander, sets the data, the template path, and renders the remplate
func htmlExample() (string, Handler) {
	uri := "/htmlexample"
	handler := func(w http.ResponseWriter, r *http.Request) {
		templateData := ExampleHtmlData{
			"This is the Example Heading",
			"Here is the example description for the paragraph in the template.",
			[]string{"Learn", "Go", "Have", "Fun", "Grow"},
		}

		templatePath := "templates/htmlexample.go.html"
		renderHtmlTemplate(w, templatePath, templateData)
	}
	return uri, handler
}

// renderHtmlTemplate - Renders the template using the http response writer
func renderHtmlTemplate(w http.ResponseWriter, templatePath string, templateData ExampleHtmlData) {
	w.Header().Set("Content-Type", "text/html")
	htmlTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}

	err = htmlTemplate.Execute(w, templateData)

	if err != nil {
		panic(err)
	}
}
