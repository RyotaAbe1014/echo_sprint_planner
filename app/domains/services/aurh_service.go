package services

import (
	"echo_sprint_planner/app/domains/repositories"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
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
	// クレームの設定
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // 有効期限

	// トークンの作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// シークレットキーでサイン
	secret := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
