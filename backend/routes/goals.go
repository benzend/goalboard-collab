package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/benzend/goalboard/utils/"
	"github.com/benzend/goalboard/database"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
)

type setGoal struct {
	Name           string `json:"Name"`
	LongTermTarget string `json:"LongTermTarget"`
	TargetPerDay   string `json:"TargetPerDay"`
}

var hmacSampleSecret = []byte("secrect")


func Goals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	HelperMethods.enableCors(&w)

	var body setGoal

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Println("failed to connect to database", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Construct the query
	query := "INSERT INTO goals_ (Name, TargetPerDay, LongTermTarget, user_id) VALUES ($1, $2, $3, $4)"

	// Parse and validate the JWT token from the cookie
	sessionInfo, err := req.Cookie("jwt_token")
	if err != nil {
		// No session cookie found
		http.Error(w, "No session cookie found", http.StatusUnauthorized)
		return
	}

	tokenString := sessionInfo.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the byte array representation of the secret key
		return hmacSampleSecret, nil
	})
	if err != nil {
		log.Println("failed to parse token", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		// Token is invalid or claims couldn't be extracted
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	var userID int

	// Query to fetch the user ID based on the username obtained from the token
	getUserQuery := "SELECT id FROM user_ WHERE username = $1;"

	// Extract the username from the JWT token claims
	username, ok := claims["username"]
	if !ok {
		// Username not found or not in expected format
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	// Query the database to get the user ID by username
	err = db.QueryRow(getUserQuery, username).Scan(&userID)
	if err != nil {
		log.Println("failed to get user", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// Execute the query with the retrieved user ID
	_, err = db.Exec(query, body.Name, body.TargetPerDay, body.LongTermTarget, userID)
	if err != nil {
		log.Println("failed to insert goal data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// If everything is fine, send a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Goal data inserted successfully")
}
