package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/benzend/goalboard/backend/models"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Goals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	switch req.Method {
	case http.MethodGet:
		
	case http.MethodPost:
		var setGoal models.Goal
		err := json.NewDecoder(req.Body).Decode(&setGoal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert the goal data into the database
		err = models.InsertGoalData(ctx, &setGoal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal("Success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)

	default:
		return
	}
}