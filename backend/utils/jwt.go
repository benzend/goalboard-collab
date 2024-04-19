package utils

import (
	"fmt"

	"github.com/benzend/goalboard/env"
)

func GetJwtSecret() []byte {
	env, err := env.ReadFile(".env")
	if err != nil {
		fmt.Println("Failed to read .env file:", err)
		return nil
	}

	hmacSecret := Invariant(env["hmacSampleSecret"], "missing `hmacSampleSecret` env variable")

	return []byte(hmacSecret)
}
