package models

//// Create my own goals /CreateGoal Endpoint
import (
	"encoding/json"
	"net/http"
	"time"
)

type Goal struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Target           uint     `json:"target"`
	TargetPer        string   `json:"targetper"`       // day | week | month
	CreatedAtDate    string   `json:"datetime"`        // datetime
	UpdatedAt        string   `json:"updatetime"`      // datetime
	GoalProgress     int      `json:"progress"`        // 100% 50% ect
	ActivitesPerGoal []string `json:"ActivityPerGoal"` //will retuern a array of activies tied to the relation ship of the goal id
}

func (g *Goal) CreatedDateTime() string {
	dt := time.Now()
	return dt.Format("01-02-2006 15:04:05")
}

func (g *Goal) CreateUserGoals(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case http.MethodPost:

		Id := req.PostFormValue("Id")
		Name := req.PostFormValue("username")
		Target := req.PostFormValue("target")
		TargetPer := req.PostFormValue("targetPer")
		GoalProgress := req.PostFormValue("percentage")
		ActivitesPerGoal := req.PostFormValue("ActviityList")

		values := map[string]string{
			"Name":             Name,
			"Id":               Id,
			"Target":           Target,
			"TargetPer":        TargetPer,
			"GoalProgres":      GoalProgress,
			"ActivitesPerGoal": ActivitesPerGoal,
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
