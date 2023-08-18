package services

import (
	"echo_sprint_planner/app/domains/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(email string, password string) (string, error)
	RefreshTokenCreate(refreshToken string) (string, error)
}

type authService struct {
	ur repositories.IUserRepository
}

func NewAuthService(ur repositories.IUserRepository) IAuthService {
	return &authService{ur}
}

func (as *authService) Login(email string, password string) (string, error) {
	// 1. emailからuserを取得
	user, err := as.ur.UserFindByEmail(email)
	if err != nil {
		return "", err
	}
	// 2. passwordを比較
	// bcrypt.GenerateFromPasswordでハッシュ化したパスワードはアンハッシュすることはできない
	// そのため、bcrypt.CompareHashAndPasswordを使用して比較する(入力されたパスワードと同じハッシュ関数で比較することでできる)
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	// 3. tokenを生成
	token, err := createToken(*user.ID)
	if err != nil {
		return "", err
	}
	// 4. tokenを返却
	return token, nil
}

func (as *authService) RefreshTokenCreate(refreshToken string) (string, error) {
	return "", nil
}

// func
func createToken(userID uuid.UUID) (string, error) {
	return "", nil
}
