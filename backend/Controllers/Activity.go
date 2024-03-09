package routeControllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/models"
	"github.com/benzend/goalboard/utils"
	_ "github.com/lib/pq"
)

var body ActivityProgress

var timestamp = time.Now().Format(time.RFC3339)

// Connect to the database
var db, err = database.Connect()

type ActivityProgress struct {
	Progress string `json:"Progress"`
}

func UpdateActivityProgress(ctx context.Context, w http.ResponseWriter, req *http.Request, username interface{}, user_id int) {

}

func getActivityProgress(ctx context.Context, w http.ResponseWriter, req *http.Request, username interface{}, user_id int) {

}

func SetActivityProgress(ctx context.Context, w http.ResponseWriter, req *http.Request, username interface{}, user_id int) {

	if err != nil {
		log.Println("failed to connect to database", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		//Create custom time stamp incase any thing false we can check times on log
		utils.CreateServerErrorLog("serverLog.txt", timestamp)
		return
	}

	// Query the database to get the user ID by username
	err = db.QueryRow(models.InsertActivityProgress, username).Scan(&user_id)
	if err != nil {
		log.Println("failed to get user", err)
		http.Error(w, "server error", http.StatusInternalServerError)

		//Create custom time stamp incase any thing false we can check times on log
		utils.CreateServerErrorLog("serverLog.txt", timestamp)
		return
	}

	// Execute the query with the retrieved user ID
	_, err = db.Exec(models.InsertActivityProgress, body.Progress, user_id)

	if err != nil {
		log.Println("failed to insert progress data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		//Create custom time stamp incase any thing false we can check times on log
		utils.CreateServerErrorLog("serverLog.txt", timestamp)
		return
	}

	// If everything is fine, send a success response and return useriD
	utils.CompletedRequest(w, user_id)

	defer db.Close()

}

func startActivityController(ctx context.Context, w http.ResponseWriter, req *http.Request) {

	secretRed := utils.ReadJwtSecret()

	utils.EnableCors(&w)

	err := json.NewDecoder(req.Body).Decode(body)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		//Create custom time stamp incase any thing false we can check times on log
		utils.CreateServerErrorLog("serverLog.txt", timestamp)

		return
	}

	// Parse and validate the JWT token from the cookie
	sessionInfo, err := req.Cookie("jwt_token")
	if err != nil {
		// No session cookie found
		http.Error(w, "No session cookie found", http.StatusUnauthorized)
		return
	}

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

	if !ok {
		// Username not found or not in expected format
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	SetActivityProgress(ctx, w, req, username, userID)

}
