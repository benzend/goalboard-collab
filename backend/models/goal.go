package models

import (
	"time"
)

type Goal struct {
	ID            string
	name          string
	target        uint   // ms
	targetPer     string // day | week | month
	CREATEDATDATE string // datetime
	updatedAt     string // datetime
}

type ModelId struct {
	models.ModelId
}

func (g *Goal) Default(id string) *Goal {

	g.CREATEDATDATE := SetCreatedAtE()

	return &Goal{
		ID:        id,
		name:      "somerandomname",
		target:    60000,
		targetPer: "day",

		createdAt: g.GetCreatedAt(),
		updatedAt: time.Now().String(),
	}
}

// This will act as a method default overide since we are unable to palce default args if we wanted to test
func NewDefault() *Goal {
	goal := new(Goal)
	return goal.Default("default_id")
}
