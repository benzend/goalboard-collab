package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/models"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Goals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	type ctxKey string
	switch req.Method {
	case http.MethodGet:

	case http.MethodPost:
		var setGoal models.Goal
		err := json.NewDecoder(req.Body).Decode(&setGoal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var ctx = context.Background()
		// Hello world, the web server
		db, err := database.Connect()

		if err != nil {
			panic(err)
		}

		// put the db value into the context to be used in fns
		ctx = context.WithValue(ctx, ctxKey("db"), db)

		defer db.Close()

		sessionInfo, err := req.Cookie("jwt_token")

		if err != nil {
			if err == http.ErrNoCookie {
				// No session cookie found
				fmt.Fprintf(w, "No session cookie found")
				return
			}
			// Other error occurred
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Insert the goal data into the database
		err = models.InsertGoalData(ctx, db, &setGoal, sessionInfo.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload, err := json.Marshal("Success")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return success response
		w.WriteHeader(http.StatusOK)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(payload)

	default:
		return
	}
}
