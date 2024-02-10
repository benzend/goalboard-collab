package models

import (
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

	SetUpdatedAtDate(string)
	GetUpdatedDate() string
}

// SetId sets the ID of MyGoal.
func (g *Goal) SetGoalID(id int) {
	g.ID = id // Access id field of embedded Goal
}

// GetId returns the ID of MyGoal.
func (g *Goal) GetGoalId() int {
	return g.ID
}

func (g *Goal) SetCreatedAt() {
	g.CreatedAtDate = time.Now().String()
}

func (g *Goal) GetCreatedAt() string {
	return g.CreatedAtDate
}

func (g *Goal) SetGoalTarget(target uint) {
	g.Target = target
}

func (g *Goal) setTargetPer(setTargetPer string) {
	g.TargetPer = setTargetPer
}

func (g *Goal) getTargetPer() string {
	return g.TargetPer
}

func (g *Goal) GetGoalTarget() uint {
	return g.Target
}

func (g *Goal) SetName(name string) {
	g.Name = name
}

func (g *Goal) GetName() string {
	return g.Name
}

func (g *Goal) SetUpdatedAtDate(updatedTime string) {
	g.UpdatedAt = updatedTime
}

func (g *Goal) GetUpdatedDateTime() string {
	return g.UpdatedAt
}

// ID            int    `json:"id"`
// Name          string `json:"name"`
// Target        uint   `json:"target"`
// TargetPer     string `json:"targetper"`  // day | week | month
// CreatedAtDate string `json:"datetime"`   // datetime
// UpdatedAt     string `json:"updatetime"` // datetime
func (g *Goal) Default(id int, name string, target uint, TargetPer string, updatedTime string) *Goal {
	g.SetGoalID(id)
	g.SetName(name)
	g.SetGoalTarget(target)
	g.setTargetPer(TargetPer)
	g.SetCreatedAt()
	g.SetUpdatedAtDate(updatedTime)
	// for _, n := range setVals {

	// }

	return &Goal{
		ID:               g.GetGoalId(),
		Name:             g.GetName(),
		CreatedAtDate:    g.GetCreatedAt(),
		SetUpdatedAtDate: g.GetUpdatedDateTime(),
	}
}
