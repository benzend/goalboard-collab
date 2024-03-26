package utils

import (
	"log"
	"os"
)

func CreateServerErrorLog(filename string, customLogVal string) {
	// Create or open the file for writing. If the file doesn't exist, create it, or truncate the file if it exists.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatalf("Failed to open or create the file: %v", err)
		return
	}

	// Create a new log.Logger that writes to the file, with a custom prefix and flags.
	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Get the custom error message. Assuming cliUserInstructions is a function that returns a string.
	customErrorMsg := customLogVal

	// Log the custom error message.
	logger.Println(customErrorMsg)

	// Ensure the file is closed when the function exits.
	defer file.Close()
}
