package utils

import (
	"fmt"

	"github.com/benzend/goalboard/env"
)

func GetJwtSecret() []byte {
	env, err := env.ReadFile(".env")

	if err != nil {
		fmt.Print("failed to read jwt env")
		return nil
	}

	jwtSalt := Invariant(env["jwtsalt"], "missing `jwtsalt` env variable")

	return []byte(jwtSalt)
}
