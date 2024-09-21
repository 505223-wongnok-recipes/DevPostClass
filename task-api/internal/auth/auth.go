package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string, secret string) (string, error) {
	var t *jwt.Token

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Audience:  jwt.ClaimStrings{username},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Second)),
	})

	signedToken, err := t.SignedString([]byte(secret))
	if err != nil {
		log.Println("error signing key")
		return signedToken, err
	}

	return signedToken, nil
}
