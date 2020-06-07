package main

import (
	"fmt"
	"mangafox/scheduler"
	"os"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func main() {
	redisURL := os.Getenv("REDIS_URI")

	if redisURL == "" {
		logrus.Fatalln(fmt.Errorf("redis url not found"))
	}

	options := asynq.RedisClientOpt{Addr: redisURL}
	queue := asynq.NewClient(options)

	scheduler.InitilizeCron(queue)
	scheduler.InitilizeServer(queue)
}
