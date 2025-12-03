package models

import (
	"errors"
	"regexp"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Validate() error {
	u.Name = strings.TrimSpace(u.Name)
	if u.Name == "" {
		return errors.New("name is required")
	}
	u.Email = strings.TrimSpace(u.Email)
	if u.Email == "" {
		return errors.New("email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}
	return nil
}
