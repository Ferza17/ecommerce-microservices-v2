package dto

import "errors"

type UserLogoutByTokenRequest struct {
	Token string `json:"token"`
}

func (u *UserLogoutByTokenRequest) Validate() error {
	if u.Token == "" {
		return errors.New("token is required")
	}
	return nil
}
