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
		validation.Field(&u.Name, validation.Required.Error("名前は必須です。"), validation.Length(3, 50).Error("名前は3文字以上50文字以下です。")),
		validation.Field(&u.Email, validation.Required.Error("メールアドレスは必須です。"), is.Email.Error("メールアドレスの形式が正しくありません。")),
		validation.Field(&u.Password, validation.Required.Error("パスワードは必須です。"), validation.Length(6, 50).Error("パスワードは6文字以上50文字以下です。")),
		validation.Field(&u.IsActive, validation.Required),
	)
}
