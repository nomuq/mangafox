package store

import (
	"mangafox/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *Store) MangaCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("manga")
}

func (store *Store) CreateManga(manga model.Manga) (*mongo.InsertOneResult, error) {
	result, err := store.MangaCollection().InsertOne(store.ctx, manga)
	return result, err
}

func (store *Store) GetMangaByMALID(slug string) (model.Manga, error) {
	var result model.Manga
	filter := bson.D{primitive.E{Key: "links.mal", Value: slug}}
	err := store.MangaCollection().FindOne(store.ctx, filter).Decode(&result)
	return result, err
}

// func (store *Store) MangaIndexes() {
// 	cursor, err := store.MangaCollection().Indexes().List(store.ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for cursor.Next(store.ctx) {
// 		index := bson.D{}
// 		cursor.Decode(&index)
// 		fmt.Println(fmt.Sprintf("index found %v", index))
// 	}
// }

func (store *Store) GetAllManga() ([]model.Manga, error) {

	var mangas []model.Manga
	cursor, err := store.MangaCollection().Find(store.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(store.ctx, &mangas); err != nil {
		return nil, err
	}

	return mangas, nil
}
