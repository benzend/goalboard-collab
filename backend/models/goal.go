package models

import (
	"time"
)

type Goal struct {
	id            string
	name          string
	target        string // ms
	targetPer     string // day | week | month
	createdAtDate string // datetime
	updatedAt     string // datetime
}

type GoalModelHelperMethods interface {
	SetGoalID(uint)
	GetGoalId() uint

	SetGoalTarget(string)
	GetGoalTarget() string

	setTargetDates(string)
	getSetTargetDates() string

	SetCreatedAt(string)
	GetCreatedAt() string
}

// SetId sets the ID of MyGoal.
func (g *Goal) SetGoalID(id string) {
	g.id = id // Access id field of embedded Goal
}

// GetId returns the ID of MyGoal.
// func (g *Goal) GetGoalId() int {
// 	return g.id
// }

func (g *Goal) SetCreatedAt() {
	g.createdAtDate = time.Now().String()
}

func (g *Goal) GetCreatedAt() string {
	return g.createdAtDate
}

func (g *Goal) Default(goalId string, name string) *Goal {

	g.SetCreatedAt()
	return &Goal{
		id:   goalId,
		name: name,
	}
}
