package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/benzend/goalboard/utils"
	_ "github.com/lib/pq"
)

func ReadEnvFile(filePath string) (map[string]string, error) {
	env := make(map[string]string)

	// Open the .env file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening .env file: %v", err)
	}
	defer file.Close()

	// Read lines from the file and parse key-value pairs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			env[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading .env file: %v", err)
	}

	return env, nil
}

func Connect() (db *sql.DB, err error) {

	//Read in .env file keys

	env, err := ReadEnvFile(".env")

	if err != nil {
		panic(fmt.Sprintf("Failed to read env file: %v", err))
	}

	if err != nil {
		panic("failed to read env file")
	}

	host := utils.Invariant[string](env["host"], "missing `host` env variable")
	password := utils.Invariant[string](env["password"], "missing `password` env variable")
	port := utils.Invariant[string](env["port"], "missing `port` env variable")
	user := utils.Invariant[string](env["user"], "missing `user` env variable")
	dbname := utils.Invariant[string](env["dbname"], "missing `dbname` env variable")

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
