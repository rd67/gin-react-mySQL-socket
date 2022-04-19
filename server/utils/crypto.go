package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashString(value string) (string, error) {

	//	Convert Password string to byte slice
	var valueBytes = []byte(value)

	//	Hash password with bcrypt's min cose
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(valueBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func HashMatch(hashedValue string, value string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
}
