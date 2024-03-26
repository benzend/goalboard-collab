package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	user_model "github.com/benzend/goalboard/models/user"
	"github.com/benzend/goalboard/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


func Login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

	// Retrieve hashed password from the database based on the provided username
	var hashedPassword string
	var userId int64 // Assuming userId is needed for other purposes
	var getUserQuery = "SELECT password, id FROM user_ WHERE username = $1;"

	err = db.QueryRow(getUserQuery, body.Username).Scan(&hashedPassword, &userId)
	if err != nil {
		log.Println("failed to get user", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the hashed password from the database
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(body.Password))
	if err != nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": body.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString(utils.GetJwtSecret())
	if err != nil {
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
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	// Return success response
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	// Return the JWT token and user data
	responseData := LoginReturnData{
		Token: tokenString,
		User: user_model.GetUser{
			ID:       userId,
			Username: body.Username,
		},
	}
	json.NewEncoder(w).Encode(responseData)
}
