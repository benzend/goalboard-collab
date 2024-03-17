package main

import (
	"context"

	"log"
	"net/http"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/router"
	"github.com/benzend/goalboard/routes"
	"github.com/benzend/goalboard/utils"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secrect")

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
	ctx = context.WithValue(ctx, utils.CTX_KEY_DB, db)

	defer db.Close()

	// var newGoal models.Goal4
	router := router.NewRouter()
	router.Ctx(ctx)

	router.Post("/register", routes.Register)
	router.Post("/login", routes.Login)
	router.Post("/logout", routes.Logout)
	// router.Get("/goals", func())

	router.Get("/healthcheck", routes.HealthCheck)

	router.Build()

	log.Println("Listening for requests at http://0.0.0.0:8000/")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
