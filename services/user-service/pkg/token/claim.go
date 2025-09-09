package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	Claim struct {
		UserId    string
		CreatedAt *time.Time
		ExpiredAt *time.Time
		jwt.StandardClaims
	}
)

// IsExpiredSoon checks if the token expires within the given duration
func (c *Claim) IsExpiredSoon(duration time.Duration) bool {
	if c.ExpiredAt == nil {
		return true
	}
	return time.Now().Add(duration).After(*c.ExpiredAt)
}

// RemainingTime returns the remaining time before the token expires
func (c *Claim) RemainingTime() time.Duration {
	if c.ExpiredAt == nil {
		return 0
	}
	remaining := c.ExpiredAt.Sub(time.Now())
	if remaining < 0 {
		return 0
	}
	return remaining
}
