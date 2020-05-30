package main

import (
	"mangafox/store"
	"mangafox/tasks"
	"mangafox/worker"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func main() {

	store := store.Store{
		URL:    "mongodb://localhost:27017",
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

	options := asynq.RedisClientOpt{Addr: "localhost:6379"}
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
