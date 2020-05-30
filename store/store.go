package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Store struct {
	URL     string
	DBName  string
	client  *mongo.Client
	db      *mongo.Database
	context context.Context
}

func (store *Store) Connect() error {
	globalContext := context.Background()

	ctx, cancel := context.WithTimeout(globalContext, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(store.URL))
	if err != nil {
		return err
	}

	store.context = globalContext
	store.client = client
	store.db = client.Database(store.DBName)
	return nil
}

func (store *Store) Ping() error {
	ctx, cancel := context.WithTimeout(store.context, 10*time.Second)
	defer cancel()
	err := store.client.Ping(ctx, readpref.Primary())
	return err
}
