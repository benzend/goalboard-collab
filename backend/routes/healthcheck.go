package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/benzend/goalboard/utils"
)

func HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	w.WriteHeader(http.StatusOK)

	type Data struct {
		Status string `json:"status"`
	}
	json.NewEncoder(w).Encode(Data{ Status: "healthy my dude" })
}
