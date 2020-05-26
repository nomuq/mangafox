package main

import (
	"context"
	"mangafox/api"
	"mangafox/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v7"
	"github.com/go-redis/redis_rate/v8"

	"github.com/satishbabariya/ratelimit"
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

	redis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	redis.Ping()
	defer redis.Close()

	err = redis.FlushDB().Err()
	if err != nil {
		logrus.Panic(err)
	}

	router := &api.Router{
		Store: str,
	}

	// limiter := rate.NewLimiter(1, 1)

	redisRateLimiter := redis_rate.NewLimiter(redis)
	limiter := &ratelimit.RateLimiter{
		Limiter: redisRateLimiter,
		Rate:    redis_rate.PerMinute(120),
	}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())

	e.Use(limiter.Limit)

	router.Routes(e)

	e.Logger.Fatal(e.Start(Port))
}
