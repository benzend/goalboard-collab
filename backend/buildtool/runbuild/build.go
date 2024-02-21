package buildtool

import (
	"fmt"
	"log"
	"os"
)

type envVars struct {
	host     string
	password string
	port     int
	user     string
	dbname   string
}

func CreateFile() {

	E := &envVars{}

	envFile, err := os.Create(".env")
	if err != nil {
		fmt.Println("Error creating .env file:", err)
		os.Exit(1)
	}

	fmt.Println("Enter Database UserName")

	_, err = fmt.Scan(
		&E.password,
		&E.dbname,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(envFile, "password=\"%s\"\n", E.password)

	if err != nil {
		fmt.Println("Error writing to .env file:", err)
		os.Exit(1)
	}

	fmt.Println("Code generation completed successfully.")
	defer envFile.Close()
}
