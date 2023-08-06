package validation

import (
	"echo_sprint_planner/app/domains/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidate interface {
	UserCreateValidate(name string, email string, isActive bool, password string) error
}

type userValidate struct {
}

// constructorを使用して、serviceの構造体を生成
func NewUserValidate() IUserValidate {
	return &userValidate{}
}

// func
func (uv *userValidate) UserCreateValidate(name string, email string, isActive bool, password string) error {
	user := &models.User{
		Name:     name,
		Email:    email,
		IsActive: isActive,
		Password: password,
	}
	if err := validation.ValidateStruct(user,
		validation.Field(&user.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.IsActive, validation.Required),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 255)),
	); err != nil {
		return err
	}
	return nil
}
