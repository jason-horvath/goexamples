package server

import (
	"fmt"
	"net/http"
)

func ServerStart() {
	http.HandleFunc("/", index)
	fmt.Println("The server is listening. Go to http://localhost:3500 in your browser")
	_ = http.ListenAndServe(":3500", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	content, err := fmt.Fprintf(w, "This is the index handler in the server package.")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes in the response content: %d", content))
}
