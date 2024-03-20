package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/benzend/goalboard/auth"
	"github.com/benzend/goalboard/utils"
	_ "github.com/lib/pq"
)

var timestamp = time.Now().Format(time.RFC3339)

type ActivityProgress struct {
	Progress string `json:"progress"`
}

func CreateActivity(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	utils.EnableCors(&w)
	user, err := auth.Authorize(ctx, w, req)

	if err != nil {
		return
	}

	db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)

	if !ok {
		http.Error(w, "failed to get db", http.StatusInternalServerError)
		return
	}

	query := "INSERT INTO activity (progress, user_id) VALUES ($1, $2)"

	type Body struct {
		Progress string `json:"progress"`
	}
	var body Body

	err = json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		http.Error(w, "couldn't parse body", http.StatusBadRequest)
		return
	}

	if _, err = db.Exec(query, body.Progress, user.ID); err != nil {
		http.Error(w, "couldn't exec query", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Goal data inserted successfully")
}
