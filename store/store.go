package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func (store Store) CreateMangaIndexes() ([]string, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		mongo.IndexModel{
			Keys: bson.M{
				"links.anilist": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mal": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mangadex": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mangareader": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mangatown": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"isPublishing": 1,
			},
		},
	}

	mangaCollection := store.db.Collection(MangaCollection)
	res, err := mangaCollection.Indexes().CreateMany(ctx, indexes)
	return res, err
}

func (store Store) CreateChapterIndexes() ([]string, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		mongo.IndexModel{
			Keys: bson.M{
				"source": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"language": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"manga": 1,
			},
		},
	}

	res, err := store.db.Collection(ChapterCollection).Indexes().CreateMany(ctx, indexes)
	return res, err
}

func (store Store) CreateMappingIndexes() ([]string, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		mongo.IndexModel{
			Keys: bson.M{
				"manga":   1,
				"chapter": 1,
				"source":  1,
			},
		},
	}

	res, err := store.db.Collection(MappingCollection).Indexes().CreateMany(ctx, indexes)
	return res, err
}
