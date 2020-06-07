package worker

import (
	"context"
	"mangafox/search"
	"mangafox/store"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	store  store.Store
	search search.Search
	server *asynq.Server
	client *asynq.Client
}

func Initilize(store store.Store, search search.Search, server *asynq.Server, client *asynq.Client) Worker {
	worker := Worker{
		store:  store,
		search: search,
		server: server,
		client: client,
	}
	return worker
}

func (worker Worker) UpdateSearchIndexes(ctx context.Context, t *asynq.Task) error {
	mangas, err := worker.store.GetAllManga()
	if err != nil {
		return err
	}

	updateID, err := worker.search.IndexAllManga(mangas)
	if err != nil {
		return err
	}
	logrus.Infoln("Updating Search Indexes", updateID)

	return nil
}
