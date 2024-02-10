package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type Goal struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Target        uint   `json:"target"`
	TargetPer     string `json:"targetper"`  // day | week | month
	CreatedAtDate string `json:"datetime"`   // datetime
	UpdatedAt     string `json:"updatetime"` // datetime
}

type GoalModelHelperMethods interface {
	SetGoalID(uint)
	GetGoalId() uint

	SetGoalTarget(uint)
	GetGoalTarget() uint

	setTargetPer(string)
	getTargetPer() string

	setTargetDates(string)
	getSetTargetDates() string

	SetCreatedAt(string)
	GetCreatedAt() string

	SetUpdatedDate(string)
	GetUpdatedDate() string

	SetName(string)
	GetName(string)
}

var (
	SetName = func(g *Goal, arg interface{}) {
		if Name, ok := arg.(string); ok {
			g.Name = Name

		}

	}

	SetGoalID = func(g *Goal, arg interface{}) {
		if Id, ok := arg.(int); ok {
			g.ID = Id
		}

	}

	SetCreatedAt = func(g *Goal, arg interface{}) {
		if CreatedAtDate, ok := arg.(string); ok {
			CreatedAtDate = time.Now().String()
			g.CreatedAtDate = CreatedAtDate

		}
	}

	SetGoalTarget = func(g *Goal, target uint) {
		g.Target = target
	}

	setTargetPer = func(g *Goal, setTargetPer string) {
		g.TargetPer = setTargetPer
	}

	SetUpdatedAtDate = func(g *Goal) {
		g.UpdatedAt = time.Now().String()
	}
)

var (
	getTargetPer = func(g *Goal) string {
		return g.TargetPer
	}

	GetGoalId = func(g *Goal) int {
		return g.ID
	}

	GetCreatedAt = func(g *Goal) string {
		return g.CreatedAtDate
	}

	GetGoalTarget = func(g *Goal) uint {
		return g.Target
	}

	GetName = func(g *Goal) string {
		return g.Name
	}

	GetUpdatedDateAt = func(g *Goal) string {
		return g.UpdatedAt
	}
)

// ID            int    `json:"id"`
// Name          string `json:"name"`
// Target        uint   `json:"target"`
// TargetPer     string `json:"targetper"`  // day | week | month
// CreatedAtDate string `json:"datetime"`   // datetime
// UpdatedAt     string `json:"updatetime"` // datetime

func (g *Goal) GetUserResponse(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the form values
	userName := req.FormValue("Name")

	switch req.Method {

	case http.MethodPost:

		data := map[string]string{
			"Name": userName,
			// target := req.Form.Get("target")
			// targetPer := req.Form.Get("targetPer")
			// createdAtDate := req.Form.Get("createdAtDate")
			// updatedAt := req.Form.Get("updatedAt")
		}

		// Unmarshal the JSON data into the map

		// Decode the JSON data from the request body into the map
		err := json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Access the values from the map

		jsonResponse, err := json.Marshal(&data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err != nil {
			firstValue := data["Name"]
			SetName(g, firstValue)
			///CALL THE REST OF THE SETERS HERE AND PASS THE VALUES

			w.Write(jsonResponse)
		}

	default:
		return
	}
}

var InitGetters = func(g *Goal) *Goal {
	return &Goal{
		ID:            GetGoalId(g),
		Name:          GetName(g),
		Target:        GetGoalTarget(g),
		TargetPer:     getTargetPer(g),
		CreatedAtDate: GetCreatedAt(g),
		UpdatedAt:     GetUpdatedDateAt(g),
	}
}
