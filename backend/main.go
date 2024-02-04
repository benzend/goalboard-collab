package main

import (
	"io"
	"log"
	"net/http"

	"github.com/benzend/goalboard/routes"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goals", routes.GoalsIndex)

	log.Println("Listing for requests at http://localhost:8000/")
	
	log.Fatal(http.ListenAndServe(":8000", nil))

}