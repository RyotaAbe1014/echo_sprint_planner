package router

import (
	"echo_sprint_planner/app/infra/handler"
	"echo_sprint_planner/app/infra/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(ah handler.IAuthHandler, uh handler.IUserHandler, sh handler.ISprintHandler) *echo.Echo {
	e := echo.New()
	// token
	e.POST("/token", ah.Login)
	// refresh token
	e.POST("/refresh", ah.Refresh)

	// user
	e.POST("/sign_up", uh.UserCreate)

	// user
	user := e.Group("/user")
	user.GET("/list", uh.GetUserList)
	user.PUT("/update", uh.UserUpdate)
	user.DELETE("/delete", uh.UserDelete)

	// sprint
	sprint := e.Group("/sprint")
	sprint.POST("/sprint", sh.SprintCreate)
	sprint.GET("/sprint_list", sh.SprintList)
	sprint.PUT("/sprint", sh.SprintUpdate)
	sprint.DELETE("/sprint", sh.SprintDelete)

	// middleware
	middleware.Middleware(e)

	// auth middleware
	middleware.AuthMiddleware(user)
	middleware.AuthMiddleware(sprint)

	return e
}
