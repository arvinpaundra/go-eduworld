package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	val, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(val)
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
