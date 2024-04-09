// package buildtool

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/benzend/goalboard/utils"
// )

// type envVars struct {
// 	host     string
// 	password string
// 	port     int
// 	user     string
// 	dbname   string
// }

// func createServerErrorLog(filename string, customLogVal string) {
// 	// Create or open the file for writing. If the file doesn't exist, create it, or truncate the file if it exists.
// 	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

// 	if err != nil {
// 		log.Fatalf("Failed to open or create the file: %v", err)
// 	}

// 	// Create a new log.Logger that writes to the file, with a custom prefix and flags.
// 	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

// 	// Get the custom error message. Assuming cliUserInstructions is a function that returns a string.
// 	customErrorMsg := utils.CliUserInstructions(customLogVal)

// 	// Log the custom error message.
// 	logger.Println(customErrorMsg)

// 	// Ensure the file is closed when the function exits.
// 	defer file.Close()

// }

// func CreateEnvFile(filename string) {

// 	E := &envVars{}

// 	envFile, err := os.Create(filename)

// 	if err != nil {
// 		fmt.Println(
// 			utils.CliUserInstructions("Error creating .env file:"), err,
// 		)
// 		os.Exit(1)
// 	}

// 	fmt.Println(utils.CliUserInstructions("Enter Database UserName"))

// 	if err != nil {
// 		fmt.Println(utils.CliUserInstructions("Enter Database UserName"), err)
// 		os.Exit(1)
// 	}

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = fmt.Fprintf(envFile, "password=\"%s\"\n", E.password)

// 	_, err = fmt.Scan(
// 		&E.password,
// 		&E.dbname,
// 	)

// 	fmt.Println(utils.CliUserInstructions("Code generation completed successfully."))
// 	defer envFile.Close()
// }
//