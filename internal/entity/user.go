package entity

import (
	"regexp"
	"strings"

	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
)
const maxEmailLength = 254

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)


type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-",omitempty`
	Role         string    `json:"role"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errorsx.ErrEmptyField
	}
	if len([]rune(u.Email)) > maxEmailLength {
		return errorsx.ErrEmailTooLong
	}
	if !emailRegex.MatchString(u.Email) {
		return errorsx.ErrInvalidEmail
	}
	parts := strings.Split(u.Email, "@")
	if len(parts) != 2 {
		return errorsx.ErrInvalidEmail
	}

	if u.Password == "" {
		return errorsx.ErrEmptyField
	}
	if u.Role == "" {
		return errorsx.ErrEmptyField
	}
	return nil
}
