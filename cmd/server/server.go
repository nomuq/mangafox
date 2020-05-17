package main

import (
	"context"
	"log"
	"mangafox/store"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	Port = ":8080"
)

func main() {
	// Formatter := new(log.TextFormatter)
	// Formatter.TimestampFormat = "02-01-2006 15:04:05"
	// Formatter.FullTimestamp = true
	// log.SetFormatter(Formatter)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	str, err := store.New(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Panic(err)
	}
	defer str.Client.Disconnect(ctx)

	str.GetAllManga()

	e.Logger.Fatal(e.Start(Port))
}
