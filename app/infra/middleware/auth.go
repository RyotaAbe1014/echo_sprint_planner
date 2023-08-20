package middleware

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(eg *echo.Group) *echo.Group {
	// 環境変数の読み込み（開発環境でのみ）
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// シークレットキーでサイン
	secret := os.Getenv("SECRET_KEY")
	// JWTの設定
	eg.Use(echojwt.JWT([]byte(secret)))
	return eg
}
