package api

import (
	"mangafox/api/manga"

	"github.com/labstack/echo"
)

type Router struct{}

func (router Router) Routes(e *echo.Echo) {
	api := e.Group("/api")

	manga := manga.Handler{}
	api.GET("/manga", manga.FindAll)
	api.GET("/manga/:id", manga.FindByID)
}
