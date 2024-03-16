package utils

import (
	"fmt"

	"github.com/benzend/goalboard/env"
)

func GetJwtSecret() []byte {
	env, err := env.ReadFile(".env")

	if err != nil {
		fmt.Print("failed to read env")
		return nil
	}

	hmacSecret := Invariant(env["hmacSampleSecret"], "missing `hmacSampleSecret` env variable")

	return []byte(hmacSecret)
}
