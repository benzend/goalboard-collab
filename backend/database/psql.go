package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// TODO: use a secret manager for the password / host / etc
const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password" // * configure this to match your psql settings
  dbname   = "goalboard_psql" // * will need to create a db beforehand until we manage this ourselves (if possible)
)

func Connect() (db *sql.DB, err error) {
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