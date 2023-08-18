package handler

import (
	"echo_sprint_planner/app/domains/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	TokenCreate(c echo.Context) error
	RefreshTokenCreate(c echo.Context) error
}

type authHandler struct {
	as services.IAuthService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// constructorを使用して、controllerの構造体を生成
func NewAuthHandler(as services.IAuthService) IAuthHandler {
	return &authHandler{as}
}

func (ah *authHandler) TokenCreate(c echo.Context) error {
	//  リクエストを構造体にバインド
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	// ログイン処理
	token, err := ah.as.Login(req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, token)
}

func (ah *authHandler) RefreshTokenCreate(c echo.Context) error {
	return nil
}
