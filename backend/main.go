package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/benzend/goalboard/backend/models"
)

// Hello world, the web server
func CreateUserGoals(w http.ResponseWriter, req *http.Request) {
	// Parse form data
	var newGoal models.Goal

	err := req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// // Access form values
	idStr := req.FormValue("id")
	id, err := strconv.Atoi(idStr)

	// target := req.Form.Get("target")
	// targetPer := req.Form.Get("targetPer")
	// createdAtDate := req.Form.Get("createdAtDate")
	// updatedAt := req.Form.Get("updatedAt")

	// Marshal the Goal object to JSON

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response

	switch req.Method {

	case http.MethodPost:

		req, err := json.Marshal(
			newGoal.Default(id),
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(req)

	default:

	}

}

func main() {
	// Hello world, the web server
	http.HandleFunc("/goals", CreateUserGoals)

	log.Println("Listening for requests at http://localhost:8000/")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
