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

func ServerStart() {
	http.HandleFunc("/", index)

	uri, handler := multiReturnHandler()
	http.HandleFunc(uri, handler)

	htmlUri, htmlHandler := htmlExample()
	http.HandleFunc(htmlUri, htmlHandler)

	fmt.Println("The server is listening. Go to http://localhost:3500 in your browser")
	_ = http.ListenAndServe(":3500", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	content, err := fmt.Fprintf(w, "This is the index handler in the server package.")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes in the response content: %d\n", content)
}

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

func htmlExample() (string, Handler) {
	uri := "/htmlexample"
	handler := func(w http.ResponseWriter, r *http.Request) {
		templateData := ExampleHtmlData{
			"This is the Example Heading",
			"Here is the example description for the paragraph in the template.",
			[]string{"Learn", "Go", "Have", "Fun", "Grow"},
		}

		templatePath := "templates/htmlexample.go.html"
		renderTemplate(w, templatePath, templateData)
	}
	return uri, handler
}

func renderTemplate(w http.ResponseWriter, templatePath string, templateData ExampleHtmlData) {
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
