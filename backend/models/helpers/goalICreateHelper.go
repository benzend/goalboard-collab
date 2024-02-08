package helpers

import (
	"time"

	"github.com/benzend/goalboard/models"
)

//WHEN GOING TO USER/GOALS IT SHOULD CREATE THE FOLLOWING INFO
/*
	ID string
	name      string //Name should come from USER HELPER class
	target    uint   // ms
	targetPer string // day | week | month
	createdAt string // datetime
	updatedAt string // datetime
*/

type ModelId interface {
	SetGoalID(uint)
	GetGoalId() uint

	SetGoalTarget(string)
	GetGoalTarget() string

	setTargetDates(string)
	getSetTargetDates() string

	SetCreatedAt(string)
	GetCreatedAt() string
}

type MyGoal struct {
	models.Goal // Embedding models.Goal
}

// SetId sets the ID of MyGoal.
func (g *MyGoal) SetId(setGoalId string) {
	g.ID = setGoalId // Access id field of embedded Goal
}

// GetId returns the ID of MyGoal.
func (g *MyGoal) GetId() string {
	return g.ID
}

func (g *MyGoal) SetCreatedAt() {
	g.CREATEDATDATE = time.Now().String()
}

func (g *MyGoal) GetCreatedAt() string {
	return g.CREATEDATDATE
}
