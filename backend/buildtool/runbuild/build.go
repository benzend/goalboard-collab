package buildtool

import (
	"fmt"
	"log"
	"os"
)

type envVars struct {
	DB_PASSWORD string
	DB_USERNAME string
}

func CreateFile() {

	E := &envVars{}

	envFile, err := os.Create(".env")
	if err != nil {
		fmt.Println("Error creating .env file:", err)
		os.Exit(1)
	}

	fmt.Println("Enter Database UserName")

	_, err = fmt.Scan(&E.DB_PASSWORD)

	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(envFile, "DB_PASSWORD=\"%s\"\n", E.DB_PASSWORD)
	if err != nil {
		fmt.Println("Error writing to .env file:", err)
		os.Exit(1)
	}

	fmt.Println("Code generation completed successfully.")
	defer envFile.Close()
}
