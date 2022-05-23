package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type ExampleHtmlData struct {
	Heading     string
	Description string
	ListItems   []string
}

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ShoeSize  int    `json:"shoe_size"`
	HasOrder  bool   `json:"has_order"`
}

type AlteredJson struct {
	Message      string `json:"message"`
	SlicePrinted string `json:"slice_printed"`
	CustomerData []Customer
}

// ServerStart - Start the server after handling the routes.
func ServerStart() {
	http.HandleFunc("/", index)

	uri, handler := multiReturnHandler()
	http.HandleFunc(uri, handler)

	htmlUri, htmlHandler := htmlExample()
	http.HandleFunc(htmlUri, htmlHandler)

	jsonUri, jsonHandler := jsonExample()
	http.HandleFunc(jsonUri, jsonHandler)

	fmt.Println("The server is listening. Go to http://localhost:3500 in your browser")
	_ = http.ListenAndServe(":3500", nil)
}

// index - The handler for the index route
func index(w http.ResponseWriter, r *http.Request) {
	content, err := fmt.Fprintf(w, "This is the index handler in the server package.")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Number of bytes in the response content: %d\n", content)
}

// multiReturnHandler - An example of using a multiple return function as a handler
func multiReturnHandler() (string, Handler) {
	uri := "/multireturnhandler"
	handler := func(w http.ResponseWriter, r *http.Request) {
		content, err := fmt.Fprintf(w, "This is the multi return handler in the server package.")

		if err != nil {
			log.Println(err)
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
		log.Println(err)
	}

	err = htmlTemplate.Execute(w, templateData)

	if err != nil {
		panic(err)
	}
}

// jsonExample - Example of handling json
func jsonExample() (string, Handler) {
	customerJson := `
    [
      {
        "first_name": "John",
        "last_name": "Doe",
        "shoe_size": 11,
        "has_order": true
      },
      {
        "first_name": "Jane",
        "last_name": "Doe",
        "shoe_size": 7,
        "has_order": false
      }
    ]
  `
	uri := "/json"

	var customerUnmarshalled []Customer
	handler := func(w http.ResponseWriter, r *http.Request) {
		err := json.Unmarshal([]byte(customerJson), &customerUnmarshalled)
		if err != nil {
			log.Println(err)
		}

		slicePrinted := fmt.Sprint(customerUnmarshalled)

		var jsonToRender AlteredJson
		jsonToRender.Message = "See the output of the slice printed, and then the altered custoemr data below."
		jsonToRender.SlicePrinted = slicePrinted

		customerUnmarshalled[0].ShoeSize = 12
		customerUnmarshalled[0].HasOrder = false
		customerUnmarshalled[1].FirstName = "Mary"
		customerUnmarshalled[1].HasOrder = true
		jsonToRender.CustomerData = customerUnmarshalled
		renderJson(w, jsonToRender)
	}

	return uri, handler
}

// renderJson - Takes the AlteredJson struct and renders it
func renderJson(w http.ResponseWriter, jsonToRender AlteredJson) {
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
