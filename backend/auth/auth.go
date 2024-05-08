package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/benzend/goalboard/utils"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func getEnv() string {
	variable, exists := os.LookupEnv("ENV")

	if !exists {
		log.Println("Failed to read ENV")
		return "development"
	}

	log.Printf("ENV=%v", variable)

	return variable
}

func Authorize(ctx context.Context, w http.ResponseWriter, req *http.Request) (user User, err error) {
	if getEnv() == "test" {
		return User{ID: 1, Username: "gwartney"}, nil
	}
	// Parse and validate the JWT token from the cookie

	sessionInfo, err := req.Cookie("jwt_token")
	if err != nil {
		http.Error(w, "no cookie", http.StatusUnauthorized)
		return
	}

	tokenString := sessionInfo.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the byte array representation of the secret key
		return []byte(utils.GetJwtSecret()), nil
	})

	if err != nil {
		log.Println("failed to parse token", err)
		http.Error(w, "failed to parse token", http.StatusInternalServerError)
		return
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		// Token is invalid or claims couldn't be extracted
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		err = fmt.Errorf("invalid token")
		return
	}

	username, ok := claims["username"]

	if !ok {
		http.Error(w, "internal server error", http.StatusUnauthorized)
		err = fmt.Errorf("no username (this shouldn't happen)")
		return
	}

	db, ok := ctx.Value(utils.CTX_KEY_DB).(*sql.DB)

	if !ok {
		http.Error(w, "failed to get db", http.StatusInternalServerError)
		err = fmt.Errorf("failed to get db")
		return
	}

	query := "SELECT id, username FROM user_ WHERE username = $1;"

	err = db.QueryRow(query, username).Scan(&user.ID, &user.Username)

	if err != nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	return
}
