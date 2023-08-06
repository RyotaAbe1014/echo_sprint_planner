package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Name     string
	Email    string
	IsActive bool
	Password string
	CreateAt *time.Time
	UpdateAt *time.Time
}

func ValidateUser(user *User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 50)),
		validation.Field(&user.IsActive, validation.Required),
	)
}
