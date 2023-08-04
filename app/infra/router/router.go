package router

import (
	"echo_sprint_planner/app/infra/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middleware.Middleware(e)
	e.GET("/", articleIndex)

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
