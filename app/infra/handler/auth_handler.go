package handler

import "github.com/labstack/echo/v4"

type IAuthHandler interface {
	TokenCreate(c echo.Context) error
	RefreshTokenCreate(c echo.Context) error
}

type authHandler struct {
}

// constructorを使用して、controllerの構造体を生成
func NewAuthHandler() IAuthHandler {
	return &authHandler{}
}

func (ah *authHandler) TokenCreate(c echo.Context) error {
	return nil
}

func (ah *authHandler) RefreshTokenCreate(c echo.Context) error {
	return nil
}
