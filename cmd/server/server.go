package main

import (
	"context"
	"mangafox/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

const (
	Port = ":8080"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	ctx := context.Background() //context.WithTimeout(context.Background(), 30*time.Second)
	str, err := store.New(ctx, "mongodb://localhost:27017")
	if err != nil {
		logrus.Panic(err)
	}
	defer str.Client.Disconnect(ctx)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(Port))
}
