package store

import (
	"context"
	"fmt"
	"mangafox/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *Store) MangaCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("manga")
}

func (store *Store) GetMangaByMangareaderID(slug string) (model.Manga, error) {
	var result model.Manga
	filter := bson.D{primitive.E{Key: "links.mangareader", Value: slug}}
	err := store.MangaCollection().FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (store *Store) CreateManga(manga model.Manga) (model.Manga, error) {
	return manga, nil
}

func (store *Store) GetAllManga() ([]model.Manga, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	database := store.Client.Database("mangafox")
	mangaCollection := database.Collection("manga")
	// chapetrsCollection := database.Collection("chapters")

	var mangas []model.Manga
	cursor, err := mangaCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &mangas); err != nil {
		panic(err)
	}
	fmt.Println(mangas)

	return mangas, nil
}
