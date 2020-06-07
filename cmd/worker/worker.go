package main

import (
	"fmt"
	"mangafox/search"
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

	search := search.Search{
		URL: "http://127.0.0.1:7700",
	}

	search.Initilize()
	err = search.CreateIndexes()
	if err != nil {
		logrus.Panicln(err)
	}

	worker := worker.Initilize(store, search, server, client)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.IndexMangadexChapter, worker.IndexMangadexChapter)
	mux.HandleFunc(tasks.UpdateSearchIndexes, worker.UpdateSearchIndexes)

	if err := server.Run(mux); err != nil {
		logrus.Fatalln(err)
	}

}
