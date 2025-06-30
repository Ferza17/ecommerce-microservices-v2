package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/pkg/errors"
	"time"
)

type (
	JWTConfig struct {
		SecretKey          string
		ExpirationDuration time.Duration
		Issuer             string
	}
)

func DefaultRefreshTokenConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:          config.Get().JwtAccessTokenSecret,
		ExpirationDuration: config.Get().JwtRefreshTokenExpirationTime,
		Issuer:             config.Get().ServiceName,
	}
}

func DefaultAccessTokenConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:          config.Get().JwtAccessTokenSecret,
		ExpirationDuration: config.Get().JwtAccessTokenExpirationTime,
		Issuer:             config.Get().ServiceName,
	}
}

func GenerateClaim(user *pb.User, config *JWTConfig) *Claim {
	now := time.Now()
	exp := now.Add(config.ExpirationDuration)
	return &Claim{
		UserId:    user.Id,
		CreatedAt: &now,
		ExpiredAt: &exp,
		StandardClaims: jwt.StandardClaims{
			Id:        user.Id,
			Subject:   user.Id,
			Audience:  user.Email,
			Issuer:    config.Issuer,
			IssuedAt:  now.Unix(),
			ExpiresAt: exp.Unix(),
			NotBefore: now.Unix(),
		},
	}
}

func GenerateToken(claim *Claim, config *JWTConfig) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateJWTToken validates and parses JWT token string
func ValidateJWTToken(tokenString string, config *JWTConfig) (*Claim, error) {
	if tokenString == "" {
		return nil, errors.New("token string cannot be empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*Claim)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Additional validation
	if err = claim.Valid(); err != nil {
		return nil, err
	}

	return claim, nil
}
