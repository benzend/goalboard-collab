package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secrect")

type JWTData struct {
	jwt.Claims
	CustomClaims map[string]string `json:"custom_claims"`
}
type User struct {
	Username string `json:"username"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Username != "gwartney" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})
	tokenString, err := token.SignedString(jwtKey)
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
	}
	http.SetCookie(w, &cookie)

	// Return success response
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the JWT cookie by setting its expiration time to a past date
	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    "",                         // Clear the value
		Expires:  time.Now().Add(-time.Hour), // Set expiration to a past time
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Hello world, the web server
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var newGoal models.Goal
	var newUser models.User

	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/logout", logoutHandler)

	http.Handle("/CreateUser", http.HandlerFunc(newUser.CreateUser))

	http.Handle("/goals", authMiddleware(http.HandlerFunc(newGoal.CreateUserGoals)))

	http.Handle("/ActivityList", authMiddleware(http.HandlerFunc(newGoal.GetActivtiesListPerGoal)))

	http.Handle("/goalprogress", authMiddleware(http.HandlerFunc(newGoal.GetGoalProgress)))

	log.Println("Listening for requests at http://0.0.0.0:8000/")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
