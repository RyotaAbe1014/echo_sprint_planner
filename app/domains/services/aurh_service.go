package services

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/domains/repositories"
	"echo_sprint_planner/app/utils"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(email string, password string) (models.Token, error)
	Refresh(refreshToken string) (models.Token, error)
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

func (as *authService) Refresh(refreshToken string) (models.Token, error) {
	// 1. リフレッシュトークンを検証
	claims, err := decodeJWT(refreshToken)
	if err != nil {
		return models.Token{}, err
	}

	// 2. トークンの有効期限を確認
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		// トークンが期限切れです
		return models.Token{}, errors.New("Token is expired")
	}

	// 3. トークンを再生成
	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return models.Token{}, errors.New("UserID is missing in refresh token")
	}

	newAccessToken, err := createAccessToken(userID)
	if err != nil {
		return models.Token{}, err
	}

	// 4. トークンを返却
	token := models.Token{
		AccessToken: newAccessToken,
	}

	return token, nil
}

// func
// シークレットキー取得

// アクセストークン生成
func createAccessToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(utils.GetEnvVar("SECRET")))
}

// リフレッシュトークン生成
func createRefreshToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(24 * 7 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(utils.GetEnvVar("SECRET")))
}

// トークン作成
func createToken(userID uuid.UUID) (models.Token, error) {
	accessToken, err := createAccessToken(userID)
	if err != nil {
		return models.Token{}, err
	}

	refreshToken, err := createRefreshToken(userID)
	if err != nil {
		return models.Token{}, err
	}

	return models.Token{
		AccessToken:  accessToken,
		RefreshToken: &refreshToken,
	}, nil
}

func decodeJWT(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnvVar("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}
