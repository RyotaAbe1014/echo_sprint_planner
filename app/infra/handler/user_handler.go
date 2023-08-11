package handler

import (
	"echo_sprint_planner/app/domains/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}

type UserUpdateRequest struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	IsActive bool      `json:"is_active"`
}
type UserDeleteRequest struct {
	ID uuid.UUID `json:"id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type IUserHandler interface {
	UserCreate(c echo.Context) error
	GetUserList(c echo.Context) error
	UserUpdate(c echo.Context) error
	UserDelete(c echo.Context) error
}

type userHandler struct {
	us services.IUserService
}

// constructorを使用して、controllerの構造体を生成
func NewUserHandler(us services.IUserService) IUserHandler {
	return &userHandler{us}
}

func (uh *userHandler) UserCreate(c echo.Context) error {
	req := new(UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	err := uh.us.UserCreate(req.Name, req.Email, req.IsActive, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success")
}

func (uh *userHandler) GetUserList(c echo.Context) error {
	userList, err := uh.us.GetUserList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, userList)
}

func (uh *userHandler) UserUpdate(c echo.Context) error {
	req := new(UserUpdateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	err := uh.us.UserUpdate(req.ID, req.Name, req.Email, req.IsActive)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success")
}

func (uh *userHandler) UserDelete(c echo.Context) error {
	req := new(UserDeleteRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	err := uh.us.UserDelete(req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, "success")
}
