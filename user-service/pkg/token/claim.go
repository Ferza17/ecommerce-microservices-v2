package token

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"time"
)

type (
	Claim struct {
		User          *pb.User
		Role          *pb.Role
		AccessControl []*pb.AccessControl
		CreatedAt     *time.Time
		ExpiredAt     *time.Time
		jwt.StandardClaims
	}
)

// HasRole checks if the user has a specific role
func (c *Claim) HasRole(roleName string) bool {
	if c.Role.String() == roleName {
		return true
	}

	return false
}

// HasRoleType checks if the user has a specific role type
func (c *Claim) HasRoleType(roleType pb.EnumRole) bool {

	if c.Role.String() == roleType.String() {
		return true
	}

	return false
}

func (c *Claim) HasAccess(serviceName, methodName string) bool {
	for _, ac := range c.AccessControl {
		if ac.ServiceName == serviceName && ac.FullMethodName == methodName {
			return true
		}
	}
	return false
}

// GetRoleID returns all role IDs for the user
func (c *Claim) GetRoleID() string {
	if c.Role != nil {
		return c.Role.Id
	}
	return ""
}

// IsAdmin checks if the user has an admin role
func (c *Claim) IsAdmin() bool {
	return c.HasRoleType(pb.EnumRole_ADMIN)
}

// IsCustomer checks if the user has an admin role
func (c *Claim) IsCustomer() bool {
	return c.HasRoleType(pb.EnumRole_CUSTOMER)
}

// IsCustomerMembership checks if the user has an admin role
func (c *Claim) IsCustomerMembership() bool {
	return c.HasRoleType(pb.EnumRole_CUSTOMER_MEMBERSHIP)
}

// GetUserID returns the user ID from the claim
func (c *Claim) GetUserID() string {
	if c.User != nil {
		return c.User.Id
	}
	return ""
}

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
