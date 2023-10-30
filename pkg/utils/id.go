package utils

import "github.com/google/uuid"

func GetID() string {
	return uuid.NewString()
}
