package models

//// Create my own goals /CreateGoal Endpoint
import (
	"context"
	"database/sql"

	"time"

	_ "github.com/lib/pq"
)

type Goal struct {
	// unique id
	ID             string `json:"id"`
	Name           string `json:"name"`
	LongTermTarget string `json:"target"`
	TargetPerDay   string `json:"target_per"`
}

type Activity struct {
	// unique id
	ID string `json:"id"`

	// data
	Progress string `json:"progress"`

	// references
	GoalID string `json:"goal_id"`
}

// / InsertGoalData inserts the goal data into the database and returns an error if any
func InsertGoalData(ctx context.Context, db *sql.DB, g *Goal, username string) error {
	// First, let's retrieve the user's ID using their username

	var userID int
	err := db.QueryRowContext(ctx, "SELECT id FROM user_ WHERE username = $1", username).Scan(&userID)

	if err != nil {
		return err
	}

	// Now that we have the userID, we can insert the goal data into the database
	query := "INSERT INTO goals_ (Name, TargetPerDay, LongTermTarget, user_id) VALUES ($1, $2, $3, $4)"
	_, err = db.ExecContext(ctx, query, g.Name, g.TargetPerDay, g.LongTermTarget, userID)
	if err != nil {
		return err
	}

	return nil
}

func (g *Goal) CreatedDateTime() string {
	dt := time.Now()
	return dt.Format("01-02-2006 15:04:05")
}

// func (g *Goal) GetActivtiesListPerGoal(w http.ResponseWriter, req *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	enableCors(&w)

// 	switch req.Method {

// 	case http.MethodGet:

// 		// Simulate activities data (replace this with your actual data) Note this would just come from the database it sefl
// 		activities := []string{"Activity 1", "Activity 2", "Activity 3"}

// 		jsonResponse, err := json.Marshal(activities)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(jsonResponse)
// 	default:
// 		jsonResponse, err := json.Marshal("Unable to find values")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		fmt.Println(string(jsonResponse))
// 		return
// 	}

// }

// func (g *Goal) GetGoalProgress(w http.ResponseWriter, req *http.Request) {

// 	enableCors(&w)

// 	w.Header().Set("Content-Type", "application/json")

// 	switch req.Method {

// 	case http.MethodGet:

// 		//Simulate the goal progress as this will be pulled from relatinal database info
// 		activities := map[string]string{
// 			"Goal1": "100%",
// 			"Goal2": "50%",
// 			"Goal3": "70%",
// 		}

// 		jsonResponse, err := json.Marshal(activities)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(jsonResponse)
// 	default:
// 		jsonResponse, err := json.Marshal("Unable to find values")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(jsonResponse)
// 		return
// 	}

//TO DO ADD IN UPDATE AND DELETE REQUEST ONCE DATABASE HOOKED UP
