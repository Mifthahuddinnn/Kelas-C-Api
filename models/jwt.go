package models

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var JWTSecret = []byte("secret")

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	claims := &TokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return signedToken, nil
}
