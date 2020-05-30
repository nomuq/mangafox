package worker

import (
	"github.com/hibiken/asynq"
	"mangafox/store"
)

type Worker struct {
	store  store.Store
	server *asynq.Server
	client *asynq.Client
}

func Initilize(store store.Store, server *asynq.Server, client *asynq.Client) Worker {
	worker := Worker{
		store:  store,
		server: server,
		client: client,
	}
	return worker
}
