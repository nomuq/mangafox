package api

import (
	"github.com/labstack/echo"
	"mangafox/api/manga"
)

type Router struct{}

func (router Router) Routes(e *echo.Echo) {
	api := e.Group("/api")

	manga := manga.Handler{}
	api.GET("/manga", manga.FindAll)
	api.GET("/manga/:id", manga.FindByID)
}
