package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwrod string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwrod), bcrypt.DefaultCost)

	if err != nil {
		panic("Error in hashing")
	}

	return string(hash)
}