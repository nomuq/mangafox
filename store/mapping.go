package store

import (
	"context"
	"mangafox/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (store *Store) MappingsCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("mapping")
}

func (store *Store) CreateMapping(mapping model.Mapping) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"$and": []bson.M{
			{"language": mapping.Language},
			{"slug": mapping.Slug},
			{"source": mapping.Source},
		},
	}
	result, err := store.MappingsCollection().UpdateOne(context.TODO(), filter, bson.M{"$set": mapping}, opts)
	return result, err
}
