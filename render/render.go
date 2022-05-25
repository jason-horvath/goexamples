package render

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jason-horvath/goexamples/schema"
)

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
