package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"log"
	"net/http"
	"time"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/routes"
	"github.com/benzend/goalboard/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secrect")

type ctxKey string

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func enableCors(w *http.ResponseWriter) {
	log.Println("setting up cors")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Hash user password from body
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func loginHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	utils.EnableCors(&w)

	var body LoginRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, ok := ctx.Value(ctxKey("db")).(*sql.DB)

	if !ok {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// Retrieve hashed password from the database based on the provided username
	var hashedPassword string
	var userId string // Assuming userId is needed for other purposes
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
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	// Return success response
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	// Return the JWT token and user data
	responseData := LoginReturnData{
		Token: tokenString,
		User: User{
			Id:       userId,
			Username: body.Username,
		},
	}
	json.NewEncoder(w).Encode(responseData)
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
		// Retrieve JWT token from the cookie
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		// Parse JWT token
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
	var ctx = context.Background()
	// Hello world, the web server
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	// put the db value into the context to be used in fns
	ctx = context.WithValue(ctx, ctxKey("db"), db)

	defer db.Close()

	// var newGoal models.Goal

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		var body LoginRequestBody

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db, ok := ctx.Value(ctxKey("db")).(*sql.DB)

		if !ok {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		password := body.Password
		hash, _ := HashPassword(password) // ignore error for the sake of simplicity

		var query = "INSERT INTO user_ (username, password) VALUES ($1, $2)"

		log.Println("inserting user...")

		if _, err := db.Exec(query, body.Username, hash); err != nil {
			log.Println("failed to insert user", err)

			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		var getUserQuery = "SELECT (username, id) FROM user_ WHERE user_.username = ?"

		var res = User{}
		if err := db.QueryRow(getUserQuery, body.Username).Scan(res); err != nil {
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
		tokenString, err := token.SignedString(jwtKey)

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
	})

	http.Handle("/login", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		loginHandler(ctx, w, req)
	}))

	http.HandleFunc("/logout", logoutHandler)

	http.Handle("/goals", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		routes.Goals(ctx, w, req)
	}))

	log.Println("Listening for requests at http://0.0.0.0:8000/")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
