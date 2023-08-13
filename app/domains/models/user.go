package models

import (
	"errors"
	"time"

	"github.com/google/uuid"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID       *uuid.UUID `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	IsActive bool       `json:"is_active"`
	Password *string
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}

func (u *User) CreateValidate() error {
	if err := nameValidate(u.Name); err != nil {
		return err
	}
	if err := emailValidate(u.Email); err != nil {
		return err
	}
	if err := passwordValidate(u.Password); err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateValidate() error {
	if err := nameValidate(u.Name); err != nil {
		return err
	}
	if err := emailValidate(u.Email); err != nil {
		return err
	}
	return nil
}

func passwordValidate(password *string) error {
	if password == nil {
		return errors.New("パスワードは必須です。")
	}
	return validation.Validate(*password, validation.Required.Error("パスワードは必須です。"), validation.Length(6, 50).Error("パスワードは6文字以上50文字以下です。"))
}

func nameValidate(name string) error {
	return validation.Validate(name, validation.Required.Error("名前は必須です。"), validation.Length(3, 50).Error("名前は3文字以上50文字以下です。"))
}

func emailValidate(email string) error {
	return validation.Validate(email, validation.Required.Error("メールアドレスは必須です。"), is.Email.Error("メールアドレスの形式が正しくありません。"))
}
