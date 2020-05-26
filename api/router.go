package api

import (
	"mangafox/api/manga"
	"mangafox/store"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Store *store.Store
}

func (router *Router) Routes(e *echo.Echo) {

	api := e.Group("/api")

	manga := &manga.MangaRouter{
		Store: router.Store,
	}

	manga.Routes(api)

}
