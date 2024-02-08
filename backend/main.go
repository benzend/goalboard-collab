package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benzend/goalboard/backend/models"
)

// Hello world, the web server
func GetUserGoals(w http.ResponseWriter, req *http.Request) {
	// Parse form data
	var newGoal models.Goal

	err := req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// // Access form values
	id := req.Form.Get("id")
	name := req.Form.Get("name")
	// target := req.Form.Get("target")
	// targetPer := req.Form.Get("targetPer")
	// createdAtDate := req.Form.Get("createdAtDate")
	// updatedAt := req.Form.Get("updatedAt")

	fmt.Printf("THE ID IS", newGoal.Default(id, name))

}

func main() {
	// Hello world, the web server
	http.HandleFunc("/goals", GetUserGoals)

	log.Println("Listening for requests at http://localhost:8000/")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
