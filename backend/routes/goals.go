package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/benzend/goalboard/database"
	"github.com/golang-jwt/jwt/v5"
)

type setGoal struct {
    Name            string `json:"Name"`
    TargetPerDay    string `json:"TargetPerDay"`
    LongTermTarget  string `json:"LongTermTarget"`
    Progress        string `json:"Progress"`
    GoalId          string    `json:"goalid"` 
}

 
var hmacSampleSecret = []byte("secrect")


func userAuthInfo(w http.ResponseWriter, req *http.Request) (jwt.MapClaims, error) {
    // Parse and validate the JWT token from the cookie
    sessionInfo, err := req.Cookie("jwt_token")
    if err != nil {
        // No session cookie found
        http.Error(w, "No session cookie found", http.StatusUnauthorized)
        return nil, err
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
        return nil, err
    }

    // Extract claims from the token
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        // Token is invalid or claims couldn't be extracted
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return nil, fmt.Errorf("Invalid token")
    }

    return claims, nil
}
func ConnectAndGetResponse(w http.ResponseWriter, req *http.Request, body *setGoal) (*sql.DB, error) {
    err := json.NewDecoder(req.Body).Decode(body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return nil, err
    }

    // Connect to the database
    db, err := database.Connect()
    if err != nil {
        log.Println("failed to connect to database", err)
        http.Error(w, "server error", http.StatusInternalServerError)
        return nil, err
    }

    return db, nil
}



func HandleError(err error, errorVal string, w http.ResponseWriter) {
	if err != nil {
		log.Println(errorVal)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Goals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	var body setGoal
 

	db,err := ConnectAndGetResponse(w,req, &body)

 
	HandleError(err, "failed to connect ", w)
	

	defer db.Close()
	claims, err := userAuthInfo(w, req)
   
	if err != nil {
        // Handle the error, such as returning an unauthorized response
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
 
	HandleError(err, "failed to get user", w)
	

	// Construct the query to insert the goal data and retrieve the goal ID
	query := "INSERT INTO goals_ (Name, TargetPerDay, LongTermTarget, user_id) VALUES ($1, $2, $3, $4) RETURNING goalid"
	// Execute the query to insert the goal data and retrieve the goal ID
	var goalID int
	
	err = db.QueryRow(query, body.Name, body.TargetPerDay, body.LongTermTarget, userID).Scan(&goalID)
 
	
	HandleError(err, "failed to insert goal data", w)

	// Construct the query to insert the progress data
	insertProgress := "INSERT INTO activity_ (goal_id, progress) VALUES ($1, $2)"
	// Execute the query to insert the progress data
	_, err = db.Exec(insertProgress, goalID, body.Progress)
	 
	HandleError(err, "Failed to insert Progress data", w)

	// If everything is fine, send a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Goal data inserted successfully")
}

func UpdateGoals(ctx context.Context, w http.ResponseWriter, req *http.Request) {
    enableCors(&w)

    var body setGoal

    // Connect to the database
    db, err := ConnectAndGetResponse(w, req, &body)
    if err != nil {
        HandleError(err, "Failed to connect", w)
        return
    }
    defer db.Close()

    // Assuming ConnectAndGetResponse fills the 'body' including 'GoalId'
    updateGoalID := body.GoalId // Corrected to use GoalId with correct capitalization

    // Update goals_ table without Progress as it does not belong to this table
    updateQuery := `
        UPDATE goals_
        SET Name = $1,
            LongTermTarget = $2,
            TargetPerDay = $3
        WHERE goalId = $4
    `

    // Execute the update query for goals_
    _, err = db.Exec(updateQuery, body.Name, body.LongTermTarget, body.TargetPerDay, updateGoalID)
    if err != nil {
        HandleError(err, "Failed to update goal data", w)
        return
    }

    // Use http.StatusOK for updates
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Goal and related activities updated successfully")
}









//Set delete golas 