package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/benzend/goalboard/env"
	"github.com/golang-jwt/jwt/v5"
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func ReadJwtSecret() []byte {
	env, err := env.ReadFile(".env")
	hmacSampleSecret := Invariant[string](env["hmacSampleSecret"], "missing `host` env variable")

	if err != nil {
		return nil
	}

	return []byte(hmacSampleSecret)

}

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
	customErrorMsg := CliUserInstructions(customLogVal)

	// Log the custom error message.
	logger.Println(customErrorMsg)

	// Ensure the file is closed when the function exits.
	defer file.Close()

}

func CliUserInstructions(userInstruction string) string {

	return userInstruction
}

func CompletedRequest(w http.ResponseWriter, userid int) int {
	w.WriteHeader(userid)
	fmt.Fprintln(w, "Goal data inserted successfully")
	return userid
}

func GetUserSessionInfo(sessionInfo *http.Cookie, secretRed []byte, w http.ResponseWriter) (jwt.MapClaims, bool) {
	tokenString := sessionInfo.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return true, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretRed, nil
	})
	if err != nil {
		log.Println("failed to parse token", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)

		return nil, true
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {

		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return nil, true
	}
	return claims, false
}
