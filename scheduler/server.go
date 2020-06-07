package scheduler

import (
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ErrorResp struct {
	Message interface{} `json:"message"`
}

func ErrorHandler(err error, c echo.Context) {

	if err, ok := err.(*echo.HTTPError); ok {
		c.JSON(err.Code, ErrorResp{
			Message: err.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResp{
		Message: err.Error(),
	})
	return
}

func InitilizeServer(queue *asynq.Client) {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Gzip())

	e.HTTPErrorHandler = ErrorHandler

	r := Router{
		queue: queue,
	}

	r.Routes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
