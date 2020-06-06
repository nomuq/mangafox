package store

import (
	"context"
	"mangafox/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MangaCollection = "manga"

func (store Store) CreateManga(manga models.Manga) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	result, err := store.db.Collection(MangaCollection).InsertOne(ctx, manga)
	if err != nil {
		return primitive.NewObjectID(), err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return objectID, err
	}

	return primitive.NewObjectID(), err
}

func (store Store) FindManga(id primitive.ObjectID) (models.Manga, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	var result models.Manga
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	err := store.db.Collection(MangaCollection).FindOne(ctx, filter).Decode(&result)
	return result, err
}

func (store Store) GetMangaByMangadexID(id string) (models.Manga, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	var result models.Manga
	filter := bson.D{primitive.E{Key: "links.mangadex", Value: id}}
	err := store.db.Collection(MangaCollection).FindOne(ctx, filter).Decode(&result)
	return result, err
}
