package dto

import "errors"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *CreateUserRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	if c.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
