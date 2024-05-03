package main

import (
	"context"

	"log"
	"net/http"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/router"
	"github.com/benzend/goalboard/routes"
)

func main() {
	var ctx = context.Background()
	// Hello world, the web server
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	// put the db value into the context to be used in fns

	defer db.Close()

	// var newGoal models.Goal4
	router := router.NewRouter()

	router.Ctx(ctx)

	router.Get("/healthcheck", routes.HealthCheck)
	router.Post("/register", routes.Register)
	router.Post("/login", routes.Login)
	router.Post("/logout", routes.Logout)
	router.Get("/goals", routes.GetGoals)
	router.Post("/goals", routes.CreateGoal)
	router.Post("/activities", routes.CreateActivity)

	router.Build()

	log.Println("Listening for requests at http://0.0.0.0:8000/")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
