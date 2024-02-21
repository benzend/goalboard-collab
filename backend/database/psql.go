package database

import (
	"database/sql"
	"fmt"

	"github.com/benzend/goalboard/backend/env"
	_ "github.com/lib/pq"
)



func Connect() (db *sql.DB, err error) {

	//Read in .env file keys

	env, err := env.ReadFile(".env")

	host := env["host"]
	password := env["password"]
	port := 5432
	user := env["user"]
	dbname := env["dbname"]

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return
}
