package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"
)

func (u *Credentials) ValidateCredentials() bool {
	var isEmailValid bool = validateEmail(u.Useremail)
	var isPasswordValid bool = validatePassword(u.Password)
	var isNameValid bool = validateUserName(u.Username)
	return isEmailValid && isPasswordValid && isNameValid

}

func validateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func validatePassword(password string) bool {
	return len(strings.TrimSpace(password)) > 3
}

func validateUserName(name string) bool {
	return len(strings.TrimSpace(name)) > 3
}

func GenerateHash(password string) (string, error) {
	// Parameters for Argon2
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Encode the hash as base64
	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	// Return the encoded hash and salt
	return fmt.Sprintf("%s:%s", saltBase64, hashBase64), nil
}

func VerifyHash(password, encodedHash string) bool {
	parts := strings.Split(encodedHash, ":")
	saltBase64, hashBase64 := parts[0], parts[1]

	salt, _ := base64.StdEncoding.DecodeString(saltBase64)
	hash, _ := base64.StdEncoding.DecodeString(hashBase64)

	// Generate a new hash from the provided password and the salt
	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Use subtle.ConstantTimeCompare to compare the new hash with the stored one
	return subtle.ConstantTimeCompare(newHash, hash) == 1
}
