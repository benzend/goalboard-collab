package routeControllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/utils"
	_ "github.com/lib/pq"
)

type setGoal struct {
	Name           string `json:"name"`
	LongTermTarget string `json:"longTermTarget"`
	TargetPerDay   string `json:"targetPerDay"`
}

func Goals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	utils.EnableCors(&w)
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

	// Parse and validate the JWT token from the cookie
	sessionInfo, err := req.Cookie("jwt_token")
	if err != nil {
		// No session cookie found
		http.Error(w, "No session cookie found", http.StatusUnauthorized)
		return
	}

	// Construct the query
	query := "INSERT INTO goals_ (Name, TargetPerDay, LongTermTarget, user_id) VALUES ($1, $2, $3, $4)"

	secretRed := utils.ReadJwtSecret()

	// Validate the signing method
	// Return the byte array representation of the secret key
	// Extract claims from the token
	// Token is invalid or claims couldn't be extracted
	claims, shouldReturn := utils.GetUserSessionInfo(sessionInfo, secretRed, w)
	if shouldReturn {
		return
	}

	var userID int

	// Extract the username from the JWT token claims
	username, ok := claims["username"]

	// Query to fetch the user ID based on the username obtained from the token
	getUserQuery := "SELECT id FROM user_ WHERE username = $1;"

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
	defer db.Close()
}
