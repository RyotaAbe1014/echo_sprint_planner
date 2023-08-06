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

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 50)),
		validation.Field(&u.IsActive, validation.Required),
	)
}
