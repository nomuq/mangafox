package scheduler

import (
	"crypto/subtle"
	"mangafox/source/mangadex"
	"net/http"
	"path"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type Router struct {
	queue *asynq.Client
}

func (router Router) Routes(e *echo.Echo) {
	api := e.Group("/api")

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("mangafox")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("mangafox")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	api.POST("/manga", router.HandleMangadexMangaSync)
	api.POST("/latest", router.HandleMangadexLatestSync)
}

func (router Router) HandleMangadexMangaSync(c echo.Context) error {
	type Params struct {
		Source string `json:"source" form:"source" query:"source"`
		ID     int64  `json:"id" form:"id" query:"id"`
	}

	params := new(Params)
	if err := c.Bind(params); err != nil {
		return err
	}

	if params.ID == 0 || params.Source == "" {
		return echo.ErrBadRequest
	}

	md := mangadex.Initilize()
	manga, err := md.GetInfo(strconv.FormatInt(params.ID, 10))
	if err != nil {
		return err
	}

	for _, chapter := range manga.Chapters {
		IndexChapter(strconv.FormatInt(params.ID, 10), strconv.FormatInt(chapter.ID, 10), router.queue)
	}

	return c.JSON(http.StatusOK, manga)
}

func (router Router) HandleMangadexLatestSync(c echo.Context) error {

	type Params struct {
		Source string `json:"source" form:"source" query:"source"`
		Token  string `json:"token" form:"token" query:"token"`
	}

	params := new(Params)
	if err := c.Bind(params); err != nil {
		logrus.Fatal(err)
	}

	if params.Token == "" || params.Source == "" {
		return echo.ErrBadRequest
	}

	md := mangadex.Initilize()
	items, err := md.Latest(params.Token)
	if err != nil {
		return err
	}

	for _, item := range items {
		chapter := path.Base(item.Link)
		manga := path.Base(item.MangaLink)
		IndexChapter(manga, chapter, router.queue)
	}

	return c.JSON(http.StatusOK, items)
}
