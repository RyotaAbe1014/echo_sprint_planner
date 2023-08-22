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
	e.POST("auth/token/", ah.Login)
	// refresh token
	e.POST("auth/refresh/", ah.Refresh)

	// user
	e.POST("user/create/", uh.UserCreate)

	// user
	user := e.Group("/user")
	middleware.AuthMiddleware(user)
	user.GET("/list/", uh.GetUserList)
	user.PUT("/update/", uh.UserUpdate)
	user.DELETE("/delete/", uh.UserDelete)

	// sprint
	sprint := e.Group("/sprint/")
	middleware.AuthMiddleware(sprint)
	sprint.POST("/create/", sh.SprintCreate)
	sprint.GET("/list/", sh.SprintList)
	sprint.PUT("/update/", sh.SprintUpdate)
	sprint.DELETE("/delete/", sh.SprintDelete)

	return e
}
