package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jason-horvath/goexamples/render"
	"github.com/jason-horvath/goexamples/schema"
)

type Handler func(w http.ResponseWriter, r *http.Request)

// ServerStart - Start the server after handling the routes.
func ServerStart() {
	http.HandleFunc("/", index)

	uri, handler := multiReturnHandler()
	http.HandleFunc(uri, handler)

	htmlUri, htmlHandler := htmlExample()
	http.HandleFunc(htmlUri, htmlHandler)

	cachedUri, cachedHandler := cachedExample()
	http.HandleFunc(cachedUri, cachedHandler)

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
		templateData := schema.ExampleHtmlData{
			Heading:     "This is the Example Heading",
			Description: "Here is the example description for the paragraph in the template.",
			ListItems:   []string{"Learn", "Go", "Have", "Fun", "Grow"},
		}

		templatePath := "./templates/static/htmlexample.go.html"
		render.HtmlTemplate(w, templatePath, templateData)
	}
	return uri, handler
}

// cachedExample - Rendering tempaltes using the built cache in the render package.
func cachedExample() (string, Handler) {
	uri := "/cachedexample"
	handler := func(w http.ResponseWriter, r *http.Request) {
		templateData := schema.ExampleHtmlData{
			Heading:     "This template is using a cache",
			Description: "Here is the cached description for the paragraph in the template.",
			ListItems:   []string{"Learn", "Go", "Have", "Fun", "Grow", "And", "Cache", "Stuff"},
		}
		pageName := "cachedexample.go.html"
		render.CachedHtmlTemplate(w, pageName, templateData)
	}
	return uri, handler
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

	var customerUnmarshalled []schema.Customer
	handler := func(w http.ResponseWriter, r *http.Request) {
		err := json.Unmarshal([]byte(customerJson), &customerUnmarshalled)
		if err != nil {
			log.Println(err)
		}

		slicePrinted := fmt.Sprint(customerUnmarshalled)

		var jsonToRender schema.OutputJason
		jsonToRender.Message = "See the output of the slice printed, and then the altered custoemr data below."
		jsonToRender.SlicePrinted = slicePrinted

		customerUnmarshalled[0].ShoeSize = 12
		customerUnmarshalled[0].HasOrder = false
		customerUnmarshalled[1].FirstName = "Mary"
		customerUnmarshalled[1].HasOrder = true
		jsonToRender.CustomerData = customerUnmarshalled
		render.JsonOutput(w, jsonToRender)
	}

	return uri, handler
}
