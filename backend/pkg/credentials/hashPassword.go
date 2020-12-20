package credentials

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//CreateHash creates a hash to be stored in the Database, simple wrapper around bcrypt
func CreateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %s", err)
	}

	return string(hash), nil
}

//CheckPassword is another thin wrapper around bcrypt to Check if the password for the user is correct
func CheckPassword(hashedPassword string, password string) (error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}