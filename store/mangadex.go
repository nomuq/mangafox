package store

import (
	"mangafox/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *Store) GetMangaByMangadexID(slug string) (model.Manga, error) {
	var result model.Manga
	filter := bson.D{primitive.E{Key: "links.mangadex", Value: slug}}
	err := store.MangaCollection().FindOne(store.Context, filter).Decode(&result)
	return result, err
}
