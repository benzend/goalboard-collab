package routes

import (
	"fmt"
	"net/http"

	"github.com/benzend/goalboard/models/helpers"
)

func GoalsIndex(w http.ResponseWriter, req *http.Request) {
	// Create a slice of MyGoal
	myGoals := []*helpers.MyGoal{}

	// Generate and set IDs for the goals
	for i := 0; i < 4; i++ {
		// Create a new MyGoal instance
		goal := &helpers.MyGoal{}

		// Set ID for the goal
		goal.SetId(fmt.Sprintf("Goal%d", i+1))

		// Append the goal to the slice
		myGoals = append(myGoals, goal)
	}

	//Map the Goals ID to the array
	mappedIDs := Map(myGoals, func(g *helpers.MyGoal) string {
		return g.GetId()
	})

	// Print the mapped IDs
	fmt.Println("Mapped IDs:", mappedIDs)
}
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
