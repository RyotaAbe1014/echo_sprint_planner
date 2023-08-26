package router

import (
	"echo_sprint_planner/app/infra/handler"
	"echo_sprint_planner/app/infra/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(ah handler.IAuthHandler, uh handler.IUserHandler, sh handler.ISprintHandler) *echo.Echo {
	e := echo.New()
	middleware.Middleware(e)
	v1 := e.Group("/api/v1")

	// auth
	v1.POST("/auth/token", ah.Login)
	// refresh token
	v1.POST("/auth/refresh", ah.Refresh)
	// get me
	v1.GET("/auth/me", ah.AuthenticatedUser)
	// user
	v1.POST("/user/create", uh.UserCreate)

	// user
	user := v1.Group("/user")
	middleware.AuthMiddleware(user)
	user.GET("/list", uh.GetUserList)
	user.PUT("/update", uh.UserUpdate)
	user.DELETE("/delete", uh.UserDelete)

	// sprint
	sprint := v1.Group("/sprint")
	middleware.AuthMiddleware(sprint)
	sprint.POST("/create", sh.SprintCreate)
	sprint.GET("/list", sh.SprintList)
	sprint.PUT("/update", sh.SprintUpdate)
	sprint.DELETE("/delete", sh.SprintDelete)

	return e
}
