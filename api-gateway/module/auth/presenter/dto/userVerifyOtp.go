package dto

import "errors"

type (
	UserVerifyOtpRequest struct {
		Otp string `json:"otp"`
	}
	UserVerifyOtpResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (c *UserVerifyOtpRequest) Validate() error {
	if c.Otp == "" {
		return errors.New("otp is required")
	}
	return nil
}
