package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//HashAndSalt password-hash
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
