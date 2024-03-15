package utils

import "github.com/benzend/goalboard/env"

func GetJwtSecret() []byte {
	env, err := env.ReadFile(".env")
	jwtSalt := Invariant(env["jwt_salt"], "missing `jwt_salt` env variable")

	if err != nil {
		return nil
	}

	return []byte(jwtSalt)
}
