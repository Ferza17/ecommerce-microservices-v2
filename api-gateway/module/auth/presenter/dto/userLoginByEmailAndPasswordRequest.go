package dto

import "errors"

type UserLoginByEmailAndPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserLoginByEmailAndPasswordRequest) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
