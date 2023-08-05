package services

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	UserCreate(name string, email string, isActive bool, password string) (nil error)
}

type userService struct {
	ur repositories.IUserRepository
}

// constructorを使用して、serviceの構造体を生成
func NewUserService(ur repositories.IUserRepository) IUserService {
	return &userService{ur}
}

// func
func (us *userService) UserCreate(name string, email string, isActive bool, password string) error {
	password, err := hashPassword(password)
	if err != nil {
		return err
	}
	user := &models.User{
		Name:     name,
		Email:    email,
		IsActive: isActive,
		Password: password,
	}
	if err := us.ur.UserCreate(user); err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
