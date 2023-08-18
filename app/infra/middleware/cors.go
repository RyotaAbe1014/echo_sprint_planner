package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func corsMiddleware(e *echo.Echo) *echo.Echo {
	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")}, // 許可するオリジン
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, // 許可するヘッダー
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken}, // CSRF対策のためのヘッダー
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"}, // 許可するHTTPメソッド
		AllowCredentials: false,                                    // cookieを使う場合はtrue
	}))
	return e
}
