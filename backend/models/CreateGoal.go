package models

//// Create my own goals /CreateGoal Endpoint
// import (
// 	_ "github.com/lib/pq"
// )

// type Activity struct {

// 	// data
// 	Progress string `json:"progress"`

// 	// references
// 	GoalID string `json:"goal_id"`
// }

// func (g *Goal) CreatedDateTime() string {
// 	dt := time.Now()
// 	return dt.Format("01-02-2006 15:04:05")
// }

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
