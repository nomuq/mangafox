package main

import (
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"mangafox/cache"
	"mangafox/store"
	"mangafox/tasks"
	"mangafox/worker"
)

func main() {
	//defer profile.Start(profile.MemProfile).Stop()

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

	cache := cache.Cache{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
	}

	worker := worker.Initilize(store, cache)

	options := asynq.RedisClientOpt{Addr: "localhost:6379"}
	server := asynq.NewServer(options, asynq.Config{
		Concurrency: 1,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.IndexMangadexChapter, worker.IndexMangadexChapter)

	if err := server.Run(mux); err != nil {
		logrus.Fatalln(err)
	}

}
