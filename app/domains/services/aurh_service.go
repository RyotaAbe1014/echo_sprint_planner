package services

import "echo_sprint_planner/app/domains/repositories"

type IAuthService interface {
	TokenCreate(email string, password string) (string, error)
	RefreshTokenCreate(refreshToken string) (string, error)
}

type authService struct {
	ur repositories.IUserRepository
}

func NewAuthService(ur repositories.IUserRepository) IAuthService {
	return &authService{ur}
}

func (as *authService) TokenCreate(email string, password string) (string, error) {
	return "", nil
}

func (as *authService) RefreshTokenCreate(refreshToken string) (string, error) {
	return "", nil
}
