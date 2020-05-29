package main

import (
	"context"
	"mangafox/service"
	"mangafox/store"
	"mangafox/tasks"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()
	store, err := store.New(ctx, "mongodb://localhost:27017")
	if err != nil {
		logrus.Fatalln(err)
	}

	service := service.Service{
		Store: store,
	}

	options := asynq.RedisClientOpt{Addr: "localhost:6379"}
	server := asynq.NewServer(options, asynq.Config{
		Concurrency: 2,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(string(tasks.IndexMangadexChapter), service.IndexMangadexChapter)

	if err := server.Run(mux); err != nil {
		logrus.Fatalln(err)
	}
}
