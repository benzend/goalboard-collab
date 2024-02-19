package buildenvcmd

import (
	"fmt"
	"os"
	"os/exec"
)

type envVars struct {
	DB_PASSWORD string
	DB_USERNAME string
}

func main(e *envVars) {
	// Define your GitHub u

	// Create or open the .env file
	envFile, err := os.Create(".env")
	if err != nil {
		fmt.Println("Error creating .env file:", err)
		os.Exit(1)
	}
	defer envFile.Close()

	// Set the environment variables
	// Write environment variables to the .env file
	_, err = fmt.Fprintf(envFile, "DB_USERNAME=%s\nDB_PASSWORD=%s\n", e.DB_PASSWORD, e.DB_PASSWORD)

	os.Setenv("DB_USERNAME", e.DB_USERNAME)
	os.Setenv("DB_PASSWORDY", e.DB_PASSWORD)

	// Run 'go generate'
	cmd := exec.Command("go", "generate")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err2 := cmd.Run()
	if err2 != nil {
		fmt.Println("Error running 'go generate':", err2)
		os.Exit(1)
	}

	fmt.Println("Code generation completed successfully.")
}
