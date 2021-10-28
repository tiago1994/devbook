package models

import (
	"devbook/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Preparate(step string) error {
	if requestError := user.validate(step); requestError != nil {
		return requestError
	}

	if requestError := user.format(step); requestError != nil {
		return requestError
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("the name field is required")
	}
	if user.Nick == "" {
		return errors.New("the nick field is required")
	}
	if user.Email == "" {
		return errors.New("the email field is required")
	}
	if requestError := checkmail.ValidateFormat(user.Email); requestError != nil {
		return errors.New("invalid email field")
	}
	if step == "create" && user.Password == "" {
		return errors.New("the password field is required")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		passwordHash, requestError := security.Hash(user.Password)
		if requestError != nil {
			return requestError
		}

		user.Password = string(passwordHash)
	}

	return nil
}
