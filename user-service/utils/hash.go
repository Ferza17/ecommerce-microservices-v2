package utils

import "golang.org/x/crypto/bcrypt"

func Hashed(password string) (response string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	if err != nil {
		return "", err
	}
	response = string(hashed)
	return
}

func Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
