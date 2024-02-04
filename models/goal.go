package models

import "time"

type Goal struct {
	id string
	name string
	target uint // ms
	targetPer string // day | week | month

	createdAt string // datetime
	updatedAt string // datetime
}

func (g Goal) Default() Goal {
	return Goal {
		id: "ksjfkjdjf",
		name: "somerandomname",
		target: 60000,
		targetPer: "day",

		createdAt: time.Now().String(),
		updatedAt: time.Now().String(),
	}
}

func (g Goal) Create() {
	g.createdAt = time.Now().String()
}

func (g Goal) Id() string {
	return g.id
}