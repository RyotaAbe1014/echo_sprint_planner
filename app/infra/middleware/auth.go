package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(e *echo.Echo) *echo.Echo {
	// JWTの設定
	e.Use(echojwt.JWT([]byte("secret")))
	return e
}
