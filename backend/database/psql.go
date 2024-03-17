package database

import (
	"database/sql"
	"fmt"

	"github.com/benzend/goalboard/env"
	"github.com/benzend/goalboard/utils"
	_ "github.com/lib/pq"
)



func Connect() (db *sql.DB, err error) {

	//Read in .env file keys

	env, err := env.ReadFile(".env")

	if err != nil {
		panic(fmt.Sprintf("Failed to read env file: %v", err))
	}

	host := utils.Invariant(env["host"], "missing `host` env variable")
	password := utils.Invariant(env["password"], "missing `password` env variable")
	port := utils.Invariant(env["port"], "missing `port` env variable")
	user := utils.Invariant(env["user"], "missing `user` env variable")
	dbname := utils.Invariant(env["dbname"], "missing `dbname` env variable")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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
