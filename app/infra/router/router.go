package router

import (
	"echo_sprint_planner/app/infra/handler"
	"echo_sprint_planner/app/infra/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(ah handler.IAuthHandler, uh handler.IUserHandler, sh handler.ISprintHandler) *echo.Echo {
	e := echo.New()
	// middleware
	middleware.Middleware(e)

	// token
	e.POST("/token", ah.Login)
	// refresh token
	e.POST("/refresh", ah.Refresh)

	// user
	e.POST("/user", uh.UserCreate)
	e.GET("/user_list", uh.GetUserList)
	e.PUT("/user", uh.UserUpdate)
	e.DELETE("/user", uh.UserDelete)

	// sprint
	e.POST("/sprint", sh.SprintCreate)
	e.GET("/sprint_list", sh.SprintList)
	e.PUT("/sprint", sh.SprintUpdate)
	e.DELETE("/sprint", sh.SprintDelete)

	return e
}
