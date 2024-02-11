package models

//// Create my own goals /CreateGoal Endpoint
import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Goal struct {
	ID           string   `json:"Id"`
	Name         string   `json:"username"`
	Target       string   `json:"target"`
	TargetPer    string   `json:"targetPer"`
	GoalProgress string   `json:"percentage"`
	Activities   []string `json:"ActivityList"`
}

func (g *Goal) CreatedDateTime() string {
	dt := time.Now()
	return dt.Format("01-02-2006 15:04:05")
}

func (g *Goal) CreateUserGoals(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case http.MethodPost:

		err := json.NewDecoder(req.Body).Decode(&g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		values := map[string]string{
			"Name":             g.Name,
			"Id":               g.ID,
			"Target":           g.Target,
			"TargetPer":        g.TargetPer,
			"GoalProgres":      g.GoalProgress,
			"ActivitesPerGoal": strings.Join(g.Activities, ", "),
			"CreatedAtDate":    g.CreatedDateTime(),
			"UpdatedAt":        g.CreatedDateTime(),
		}

		jsonResponse, err := json.Marshal(values)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)

	default:
		//Replace with better error for end user
		jsonResponse, err := json.Marshal("Unable to find values")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
		return
	}

}

func (g *Goal) GetActivtiesListPerGoal(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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
		w.Write(jsonResponse)
		return
	}

}

func (g *Goal) GetGoalProgress(w http.ResponseWriter, req *http.Request) {

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
