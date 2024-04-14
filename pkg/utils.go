package utils

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func HashPassword(password string) []byte {
	hashPassword := sha512.Sum512([]byte(password))
	passwordByteSlice := hashPassword[:]
	return passwordByteSlice
}

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}

func ValidateImage(filename string) error {
	allowedExtensions := []string{".png", ".jpg", ".jpeg", ".webp"}

	matchFound := false
	for _, allowedExtension := range allowedExtensions {
		if filename == allowedExtension {
			matchFound = true
			break
		}
	}

	if !matchFound {
		return fmt.Errorf("invalid file extension: %s", filename)
	}

	return nil
}

const (
	DefaultOffset = 0
	DefaultLimit  = 10
	MaxRetries    = 3
)
