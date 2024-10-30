package loginhandler

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

func GenerateJWTToken(userEmail string) (string, error) {
	claims := jwt.MapClaims{"userEmail": userEmail, "exp": time.Now().Add(time.Hour * 24).Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)

}
