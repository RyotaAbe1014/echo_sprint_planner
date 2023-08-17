package handler

import (
	"echo_sprint_planner/app/domains/services"

	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	TokenCreate(c echo.Context) error
	RefreshTokenCreate(c echo.Context) error
}

type authHandler struct {
	as services.IAuthService
}

// constructorを使用して、controllerの構造体を生成
func NewAuthHandler(as services.IAuthService) IAuthHandler {
	return &authHandler{as}
}

func (ah *authHandler) TokenCreate(c echo.Context) error {
	return nil
}

func (ah *authHandler) RefreshTokenCreate(c echo.Context) error {
	return nil
}
