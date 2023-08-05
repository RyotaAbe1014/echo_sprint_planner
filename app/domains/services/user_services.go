package services

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"
)

type IUserService interface {
	UserCreate(name string, email string, isActive bool) (nil error)
}

type userService struct {
	ur repositories.IUserRepository
}

// constructorを使用して、serviceの構造体を生成
func NewUserService(ur repositories.IUserRepository) IUserService {
	return &userService{ur}
}

// func
func (us *userService) UserCreate(name string, email string, isActive bool) error {
	user := &models.User{
		Name:     name,
		Email:    email,
		IsActive: isActive,
	}
	if err := us.ur.UserCreate(user); err != nil {
		return err
	}
	return nil
}
