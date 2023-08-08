package router

import (
	"echo_sprint_planner/app/infra/handler"
	"echo_sprint_planner/app/infra/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(uh handler.IUserHandler) *echo.Echo {
	e := echo.New()
	// middleware
	middleware.Middleware(e)

	// user
	e.POST("/user", uh.UserCreate)
	e.GET("/user_list", uh.GetUserList)

	return e
}
