package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	user_model "github.com/benzend/goalboard/models/user"
	"github.com/benzend/goalboard/pw"
	"github.com/benzend/goalboard/utils"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReturnData struct {
	Token string             `json:"token"`
	User  user_model.GetUser `json:"user"`
}

func Register(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	var body LoginRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)

	if !ok {
		http.Error(w, "server error: failed to connect db", http.StatusInternalServerError)
		log.Println("failed to grab db")
		return
	}

	password := body.Password
	hash, err := pw.HashPassword(password) // ignore error for the sake of simplicity

	if err != nil {
		http.Error(w, "server error: failed to hash password", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = user_model.Create(db, body.Username, hash)

	if err != nil {
		http.Error(w, "server error: failed to create user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	user, err := user_model.FindFromUsername(db, body.Username)

	if err != nil {
		http.Error(w, "server error: failed to find user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": body.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	log.Println("getting signed string...")
	tokenString, err := token.SignedString(utils.GetJwtSecret())

	if err != nil {
		log.Println("failed to sign string")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := map[string]any {"token": tokenString, "user": user}

	// Set the token in a cookie
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS

	}
	http.SetCookie(w, &cookie)

	// Return success response
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(payload)
}
