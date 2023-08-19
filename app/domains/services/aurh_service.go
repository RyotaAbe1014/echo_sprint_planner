package services

import (
	"echo_sprint_planner/app/domains/models"
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
	Login(email string, password string) (models.Token, error)
	RefreshTokenCreate(refreshToken string) (string, error)
}

type authService struct {
	ur repositories.IUserRepository
}

func NewAuthService(ur repositories.IUserRepository) IAuthService {
	return &authService{ur}
}

func (as *authService) Login(email string, password string) (models.Token, error) {
	// 1. emailからuserを取得
	user, err := as.ur.UserFindByEmail(email)
	if err != nil {
		return models.Token{}, err
	}
	// 2. passwordを比較
	// bcrypt.GenerateFromPasswordでハッシュ化したパスワードはアンハッシュすることはできない
	// そのため、bcrypt.CompareHashAndPasswordを使用して比較する(入力されたパスワードと同じハッシュ関数で比較することでできる)
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		return models.Token{}, err
	}
	// 3. tokenを生成
	token, err := createToken(*user.ID)
	if err != nil {
		return models.Token{}, err
	}
	// 4. tokenを返却
	return token, nil
}

func (as *authService) RefreshTokenCreate(refreshToken string) (string, error) {
	return "", nil
}

// func
func createToken(userID uuid.UUID) (models.Token, error) {
	// アクセストークンのクレーム設定
	accessClaims := jwt.MapClaims{}
	accessClaims["authorized"] = true
	accessClaims["user_id"] = userID
	accessClaims["exp"] = time.Now().Add(time.Hour * 1).Unix() // 有効期限

	// リフレッシュトークンのクレーム設定
	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = userID
	refreshClaims["exp"] = time.Now().Add(24 * 7 * time.Hour).Unix() // リフレッシュトークンの有効期限

	// トークンの作成
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// 環境変数の読み込み（開発環境でのみ）
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// シークレットキーでサイン
	secret := os.Getenv("SECRET_KEY")
	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return models.Token{}, err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return models.Token{}, err
	}

	// 構造体に格納
	tokens := models.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return tokens, nil
}
