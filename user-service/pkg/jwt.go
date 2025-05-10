package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	Claim struct {
		UserID    string
		CreatedAt *time.Time
		jwt.StandardClaims
	}
)

func GenerateToken(claim Claim, secret string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseClaimFromToken(token string, secret string) (Claim, error) {
	var claim Claim
	_, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return Claim{}, err
	}
	return claim, nil
}
