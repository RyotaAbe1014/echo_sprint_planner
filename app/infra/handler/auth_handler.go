package handler

import (
	"echo_sprint_planner/app/domains/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	Login(c echo.Context) error
	Refresh(c echo.Context) error
}

type authHandler struct {
	as services.IAuthService
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RefreshRequest struct {
	RefreshToken string `json:"refresh"`
}

// constructorを使用して、controllerの構造体を生成
func NewAuthHandler(as services.IAuthService) IAuthHandler {
	return &authHandler{as}
}

func (ah *authHandler) Login(c echo.Context) error {
	//  リクエストを構造体にバインド
	req := new(loginRequest)
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

func (ah *authHandler) Refresh(c echo.Context) error {
	req := new(RefreshRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}

	token, err := ah.as.Refresh(req.RefreshToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, token)
}
