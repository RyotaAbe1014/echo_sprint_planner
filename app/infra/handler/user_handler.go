package handler

import (
	"echo_sprint_planner/app/domains/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	UserCreate(c echo.Context) error
}

type userHandler struct {
	us services.IUserService
}

// constructorを使用して、controllerの構造体を生成
func NewUserHandler(us services.IUserService) IUserHandler {
	return &userHandler{us}
}

// UserCreate is a function to create a user
func (uh *userHandler) UserCreate(c echo.Context) error {
	err := uh.us.UserCreate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success")
}
