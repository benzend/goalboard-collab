package main

import (
	"context"
	"fmt"

	"log"
	"net/http"

	"github.com/benzend/goalboard/database"
	"github.com/benzend/goalboard/router"
	"github.com/benzend/goalboard/routes"
	"github.com/benzend/goalboard/utils"
)

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

	router.Get("/healthcheck", routes.HealthCheck)
	router.Post("/register", routes.Register)
	router.Post("/login", routes.Login)
	router.Post("/logout", routes.Logout)
	router.Get("/goals", routes.GetGoals)
	router.Post("/goals", routes.CreateGoal)
	router.Post("/activities", routes.CreateActivity)
	router.Get("/current_user", routes.GetCurrentUser)

	router.Build()

	port := 8000

	log.Printf("Listening for requests at port %v", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
