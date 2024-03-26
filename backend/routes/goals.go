package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/benzend/goalboard/auth"
	goal_model "github.com/benzend/goalboard/models/goal"
	"github.com/benzend/goalboard/utils"
	_ "github.com/lib/pq"
)

type setGoal struct {
	GoalID         string `json:"id"`
	Name           string `json:"name"`
	Progress       string `json:"progress"`
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

	goalID, err := goal_model.Create(db, body.Name, body.TargetPerDay, body.LongTermTarget, user.ID)

	if err != nil {
		log.Println("failed to insert goal data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// If everything is fine, send a success response
	type Data struct {
		GoalID int64 `json:"goal_id"`
	}
	err = json.NewEncoder(w).Encode(Data{ GoalID: goalID })

	if err != nil {
		log.Println("failed to encode json data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetGoals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	utils.EnableCors(&w)

	user, err := auth.Authorize(ctx, w, req)

	if err != nil {
		return
	}

	db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)

	if !ok {
		log.Println("failed to retrieve database")
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	query := "SELECT id, name, target_per_day, long_term_target FROM goal WHERE user_id = $1"

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

func UpdateGoals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
    utils.EnableCors(&w)

    var body setGoal

    db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)
    if !ok {
				log.Println("failed to retrieve database")
				http.Error(w, "server error", http.StatusInternalServerError)
        return
    }

    // Assuming ConnectAndGetResponse fills the 'body' including 'GoalId'
    updateGoalID := body.GoalID

    // Update goals_ table without Progress as it does not belong to this table
    updateQuery := `
        UPDATE goal
        SET name = $1,
            long_term_target = $2,
            target_per_day = $3
        WHERE id = $4
    `

    updateProgQuery := `
        UPDATE activity
        SET progress = $1
        WHERE goal_id  = $2
    `
    // Execute the update query for goals_
		_, err := db.Exec(updateQuery, body.Name, body.LongTermTarget, body.TargetPerDay, updateGoalID)
    if err != nil {
				log.Println("failed to retrieve database")
				http.Error(w, "server error", http.StatusInternalServerError)
        return
    }

		_, err = db.Exec(updateProgQuery, body.Progress, updateGoalID)
		if err != nil {
			log.Println("Error updating activity:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

    // Use http.StatusOK for updates
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Goal and related activities updated successfully")
}

func DeleteGoalAndActivities(ctx context.Context, w http.ResponseWriter, req *http.Request) {
    utils.EnableCors(&w)

    var body setGoal

    db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)
    if !ok {
				log.Println("failed to retrieve database")
				http.Error(w, "server error", http.StatusInternalServerError)
        return
    }

    updateGoalID := body.GoalID

    // Define the delete query
    deleteQuery := `
        DELETE FROM goal
        WHERE id = $1
    `

    // Execute the delete query for goals_
		_, err := db.Exec(deleteQuery, updateGoalID)
    if err != nil {
        log.Println("Failed to delete goal:", err)
        http.Error(w, "Failed to delete goal", http.StatusInternalServerError)
        return
    }

    // Send a success response
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Goal and related activities updated successfully")
}
