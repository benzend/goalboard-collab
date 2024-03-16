package user_model

import (
	"database/sql"
	"fmt"
	"log"
)

type GetUser struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
}

func Find(db *sql.DB, id int64) (user GetUser, err error) {
	log.Println("finding user...")

	query := "SELECT username, id FROM user_ WHERE id = $1"

	if err = db.QueryRow(query, id).Scan(&user.ID, &user.Username); err != nil {
		err = fmt.Errorf("find user:%v", err)
		return
	}

	return
}

func FindFromUsername(db *sql.DB, username string) (user GetUser, err error) {
	log.Println("finding user...")

	query := "SELECT username, id FROM user_ WHERE username = $1"

	if err = db.QueryRow(query, username).Scan(&user.Username, &user.ID); err != nil {
		err = fmt.Errorf("find user:%v", err)
		return
	}

	return
}

func Create(db *sql.DB, username string, password string) (err error) {
	log.Println("inserting user...")

	query := "INSERT INTO user_ (username, password) VALUES ($1, $2)"

	_, err = db.Exec(query, username, password)

	if err != nil {
		err = fmt.Errorf("add user:%v", err)
		return
	}

	return
}
