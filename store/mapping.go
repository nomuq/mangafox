package store

import (
	"context"
	"mangafox/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MappingCollection = "mapping"

func (store Store) CreateMapping(mapping models.Mapping) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	result, err := store.db.Collection(MappingCollection).InsertOne(ctx, mapping)
	if err != nil {
		return primitive.NewObjectID(), err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return objectID, err
	}

	return primitive.NewObjectID(), err
}

func (store Store) FindChapterMapping(source string, manga string, chapter string) (models.Mapping, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	var result models.Mapping
	filter := bson.M{
		"source":  source,
		"manga":   manga,
		"chapter": chapter,
	}
	err := store.db.Collection(MappingCollection).FindOne(ctx, filter).Decode(&result)
	return result, err
}
