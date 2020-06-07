package main

import (
	"fmt"
	"mangafox/store"
	"mangafox/tasks"
	"mangafox/worker"
	"os"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func main() {
	mongoURL := os.Getenv("MONGO_URI")
	redisURL := os.Getenv("REDIS_URI")

	if mongoURL == "" {
		logrus.Fatalln(fmt.Errorf("mongodb url not found"))
	}

	if redisURL == "" {
		logrus.Fatalln(fmt.Errorf("redis url not found"))
	}

	store := store.Store{
		URL:    mongoURL,
		DBName: "mangafox",
	}
	// defer client.Disconnect(ctx)

	err := store.Connect()
	if err != nil {
		logrus.Panicln(err)
	}

	err = store.Ping()
	if err != nil {
		logrus.Panicln(err)
	}

	options := asynq.RedisClientOpt{Addr: redisURL}
	server := asynq.NewServer(options, asynq.Config{
		Concurrency: 1,
	})

	client := asynq.NewClient(options)

	worker := worker.Initilize(store, server, client)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.IndexMangadexChapter, worker.IndexMangadexChapter)

	if err := server.Run(mux); err != nil {
		logrus.Fatalln(err)
	}

}
