package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/benzend/goalboard/auth"
	user_model "github.com/benzend/goalboard/models/user"
	"github.com/benzend/goalboard/utils"
)

type GetCurrentUserReturnData struct  {
	User user_model.GetUser `json:"user"`
}

func GetCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	user, err := auth.Authorize(ctx, w, r)

	if err != nil {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	// Return the JWT token and user data
	responseData := LoginReturnData{
		User: user_model.GetUser{
			ID:       user.ID,
			Username: user.Username,
		},
	}
	json.NewEncoder(w).Encode(responseData)
}
