package models

type Activity struct {
	id string
	duration uint // ms

	goalId string

	createdAt string // datetime
	updatedAt string // datetime
}