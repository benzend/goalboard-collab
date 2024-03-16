package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benzend/goalboard/pw"
	"github.com/benzend/goalboard/utils"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReturnData struct {
	Token string `json:"token"`
	User User `json:"user"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
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
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		password := body.Password
		hash, _ := pw.HashPassword(password) // ignore error for the sake of simplicity

		var query = "INSERT INTO user_ (username, password) VALUES ($1, $2)"

		log.Println("inserting user...")

		if _, err := db.Exec(query, body.Username, hash); err != nil {
			log.Println("failed to insert user", err)

			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		var getUserQuery = "SELECT username, id FROM user_ WHERE username = $1"

		var res User
		if err := db.QueryRow(getUserQuery, body.Username).Scan(&res.Id, &res.Username); err != nil {
			log.Println("failed to get user", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		// Create a JWT token
		log.Println("getting token...")
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

		payload, err := json.Marshal(LoginReturnData{Token: tokenString, User: res})

		if err != nil {
			log.Println("failed to marshal")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
