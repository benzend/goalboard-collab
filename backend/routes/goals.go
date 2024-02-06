package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/benzend/goalboard/models"
	"github.com/benzend/goalboard/models/helpers"
)

func GoalsIndex(w http.ResponseWriter, req *http.Request) {
	goals := []models.Goal{(models.Goal{}).Default(), (models.Goal{}).Default(), (models.Goal{}).Default()}

	msg := fmt.Sprintf("All of my goals: %v", Map(goals, helpers.IdMapper))
	io.WriteString(w, msg)

}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
