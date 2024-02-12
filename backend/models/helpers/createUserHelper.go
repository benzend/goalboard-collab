package helpers

//WHEN GOING TO USER/GOALS IT SHOULD CREATE THE FOLLOWING INFO
//Or WHEN A USER GOES TO /USERS/Accoutn SHOULD BE ABLE TO PULL USER DATA

/*
	ID string
	name      string //Name should come from USER HELPER class
	target    uint   // ms
	targetPer string // day | week | month
	createdAt string // datetime
	updatedAt string // datetime
*/

type ModelUser interface {
	SetGoalID(uint)
	GetGoalId() uint

	SetGoalTarget(string)
	GetGoalTarget() string

	setTargetDates(string)
	getSetTargetDates() string

	SetCreatedAt(string)
	GetCreatedAt() string
}
