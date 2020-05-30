package worker

import (
	"mangafox/store"
)

type Worker struct {
	store store.Store
	// cache cache.Cache
}

func Initilize(store store.Store) Worker {
	worker := Worker{
		store: store,
		// cache: cache,
	}
	return worker
}
