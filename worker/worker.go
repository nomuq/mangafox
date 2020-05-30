package worker

import (
	"mangafox/cache"
	"mangafox/store"
)

type Worker struct {
	store store.Store
	cache cache.Cache
}

func Initilize(store store.Store, cache cache.Cache) Worker {
	worker := Worker{
		store: store,
		cache: cache,
	}
	return worker
}
