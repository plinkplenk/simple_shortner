package utils

import (
	"errors"
	"github.com/plinkplenk/simple_shortner/internal/api/dtos"
	"net/mail"
	"regexp"
)

type UserValidator struct {
}

func (v UserValidator) IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (v UserValidator) IsValidUsername(username string) bool {
	match, err := regexp.Match(`^[a-zA-Z0-9]{5,15}$`, []byte(username))
	if err != nil {
		return false
	}
	return match
}

func (v UserValidator) IsValidPassword(password string) bool {
	lowercase, _ := regexp.MatchString(`[a-z]`, password)
	uppercase, _ := regexp.MatchString(`[A-Z]`, password)
	digit, _ := regexp.MatchString(`[0-9]`, password)
	special, _ := regexp.MatchString(`[@$!%*?&]`, password)
	length := len(password) >= 8
	valid := lowercase && uppercase && digit && special && length
	return valid
}

func (v UserValidator) ValidateLoginData(user *dtos.UserLoginDto) (string, error) {
	if user.Email == nil && user.Username == nil {
		return "", errors.New("you must provide username or email to login")
	}
	if user.Email != nil {
		return *user.Email, nil
	}
	return *user.Username, nil
}

func (v UserValidator) ValidateUserData(user *dtos.UserCreateDto) error {
	if !v.IsValidEmail(user.Email) {
		return errors.New("invalid email")
	}
	if !v.IsValidUsername(user.Username) {
		return errors.New("invalid username. Username should contains only letters, and numbers")
	}
	if !v.IsValidPassword(user.Password) {
		return errors.New("password should contains letters in upper and lower case, numbers and special symbols")
	}
	return nil
}
