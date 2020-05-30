package manga

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Handler struct{}

func (handler Handler) FindAll(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (handler Handler) FindByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, id)
}
