package models

import (
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
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required.Error("名前は必須です。"), validation.Length(3, 50).Error("名前は3文字以上50文字以下です。")),
		validation.Field(&u.Email, validation.Required.Error("メールアドレスは必須です。"), is.Email.Error("メールアドレスの形式が正しくありません。")),
		validation.Field(&u.Password, validation.Required.Error("パスワードは必須です。"), validation.Length(6, 50).Error("パスワードは6文字以上50文字以下です。")),
		validation.Field(&u.IsActive, validation.Required),
	)
}
