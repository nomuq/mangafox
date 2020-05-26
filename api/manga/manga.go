package manga

import (
	"mangafox/store"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MangaRouter struct {
	Store *store.Store
}

func (router *MangaRouter) Routes(group *echo.Group) {
	group.GET("/manga", router.findAll)
	group.GET("/manga/:id", router.findByID)
}

func (router *MangaRouter) findAll(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (router *MangaRouter) findByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, id)
}
