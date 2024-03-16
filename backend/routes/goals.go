package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/benzend/goalboard/auth"
	"github.com/benzend/goalboard/utils"
	_ "github.com/lib/pq"
)

type setGoal struct {
	Name           string `json:"name"`
	LongTermTarget string `json:"long_term_target"`
	TargetPerDay   string `json:"target_per_day"`
}

func CreateGoal(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	utils.EnableCors(&w)

	user, err := auth.Authorize(ctx, w, req)

	if err != nil {
		return
	}

	var body setGoal

	err = json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)

	if !ok {
		log.Println("failed to retrieve database", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	query := "INSERT INTO goal (name, target_per_day, long_term_target, user_id) VALUES ($1, $2, $3, $4)"

	_, err = db.Exec(query, body.Name, body.TargetPerDay, body.LongTermTarget, user.ID)
	if err != nil {
		log.Println("failed to insert goal data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// If everything is fine, send a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Goal data inserted successfully")
}

func GetGoals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	utils.EnableCors(&w)

	user, err := auth.Authorize(ctx, w, req)

	if err != nil {
		return
	}

	db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)

	if !ok {
		log.Println("failed to retrieve database", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	query := "SELECT goal_id, name, target_per_day, long_term_target FROM goal WHERE user_id = $1"

	type Goal struct {
		ID string `json:"id"`
		Name string `json:"name"`
		TargetPerDay string `json:"target_per_day"`
		LongTermTarget string `json:"long_term_target"`
	}
	var goals []Goal

	rows, err := db.Query(query, user.ID)
	if err != nil {
		log.Println("failed to insert goal data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var row Goal
		if err := rows.Scan(&row.ID, &row.Name, &row.TargetPerDay, &row.LongTermTarget); err != nil {
			http.Error(w, "row scanning error", http.StatusInternalServerError)
			return
		} else {
			goals = append(goals, row)
		}
	}

	type Res struct {
		Goals []Goal `json:"goals"`
	}
	json.NewEncoder(w).Encode(Res{ Goals: goals })
	// If everything is fine, send a success response
	w.WriteHeader(http.StatusOK)
}
