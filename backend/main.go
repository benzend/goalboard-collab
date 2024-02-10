package main

import (
	"log"
	"net/http"

	"github.com/benzend/goalboard/backend/models"
)

func main() {
	// Hello world, the web server
	var newGoal models.Goal
	http.HandleFunc("/goals", newGoal.GetUserResponse)

	log.Println("Listening for requests at http://localhost:8000/")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
