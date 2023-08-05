package handler

import (
	"echo_sprint_planner/app/domains/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

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
	req := new(UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	err := uh.us.UserCreate(req.Name, req.Email, req.IsActive)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success")
}
