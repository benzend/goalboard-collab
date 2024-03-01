package models

//// Create my own goals /CreateGoal Endpoint
import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Goal struct {
	// unique id
	ID           string    `json:"id"`
  
	// data
	Name         string    `json:"name"`
	Target       string    `json:"target"`
	TargetPer    string    `json:"target_per"`
  }
  
  type Activity struct {
	// unique id
	ID        string       `json:"id"`
  
	// data
	Progress  string       `json:"progress"`
	
	// references
	GoalID    string       `json:"goal_id"`
  }
// InsertGoalData inserts the goal data into the database and returns an error if any
func InsertGoalData(ctx context.Context, g *Goal) error {
	type ctxKey string

	db, ok := ctx.Value(ctxKey("db")).(*sql.DB)

	var query = "INSERT INTO user_ (username, password) VALUES ($1, $2)"

	log.Println("inserting user...")

	if _, err := db.Exec(query, body.Username, hash); err != nil {
		log.Println("failed to insert user", err)

		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if !ok {
		panic("failed to get db value")
	}


	var query = "INSERT INTO goals (username, password) VALUES ($1, $2)"
	// Prepare the SQL statement
	var stmt = "INSERT INTO goals 
	(Name, 
	Target, TargetPer, 
	GoalProgress, ActivitiesPerGoal, 
	CreatedAtDate, UpdatedAt) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(g.Name, g.ID, g.Target, g.TargetPer, g.GoalProgress, strings.Join(g.Activities, ", "), g.CreatedDateTime(), g.CreatedDateTime())
	if err != nil {
		return err
	}

	return nil

}

func (g *Goal) CreatedDateTime() string {
	dt := time.Now()
	return dt.Format("01-02-2006 15:04:05")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (g *Goal) GetActivtiesListPerGoal(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	enableCors(&w)

	switch req.Method {

	case http.MethodGet:

		// Simulate activities data (replace this with your actual data) Note this would just come from the database it sefl
		activities := []string{"Activity 1", "Activity 2", "Activity 3"}

		jsonResponse, err := json.Marshal(activities)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	default:
		jsonResponse, err := json.Marshal("Unable to find values")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(jsonResponse))
		return
	}

}

func (g *Goal) GetGoalProgress(w http.ResponseWriter, req *http.Request) {

	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case http.MethodGet:

		//Simulate the goal progress as this will be pulled from relatinal database info
		activities := map[string]string{
			"Goal1": "100%",
			"Goal2": "50%",
			"Goal3": "70%",
		}

		jsonResponse, err := json.Marshal(activities)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	default:
		jsonResponse, err := json.Marshal("Unable to find values")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
		return
	}

}

//TO DO ADD IN UPDATE AND DELETE REQUEST ONCE DATABASE HOOKED UP
