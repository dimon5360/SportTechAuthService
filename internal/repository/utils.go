package repository

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	salt, isExist := os.LookupEnv("PASSWORD_HASH_SALT")
	if !isExist {
		log.Println("salt not found")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), 10)
	return string(bytes), err
}

func ValidatePassword(password, raw string) bool {
	salt, isExist := os.LookupEnv("PASSWORD_HASH_SALT")
	if !isExist {
		log.Println("salt not found")
	}
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(raw+salt)) == nil
}
