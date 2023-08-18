package middleware

import (
	"github.com/labstack/echo/v4"
)

func Middleware(e *echo.Echo) *echo.Echo {
	// CORSの設定
	corsMiddleware(e)

	// ログの設定
	logMiddleware(e)

	return e
}
