package store

import (
	"context"
	"mangafox/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func (store *Store) MangaCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("manga")
}

func (store *Store) CreateManga(manga model.Manga) (*mongo.InsertOneResult, error) {
	result, err := store.MangaCollection().InsertOne(context.TODO(), manga)
	return result, err
}

// func (store *Store) MangaIndexes() {
// 	cursor, err := store.MangaCollection().Indexes().List(context.TODO())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for cursor.Next(context.TODO()) {
// 		index := bson.D{}
// 		cursor.Decode(&index)
// 		fmt.Println(fmt.Sprintf("index found %v", index))
// 	}
// }

// func (store *Store) GetAllManga() ([]model.Manga, error) {

// 	database := store.Client.Database("mangafox")
// 	mangaCollection := database.Collection("manga")
// 	// chapetrsCollection := database.Collection("chapters")

// 	var mangas []model.Manga
// 	cursor, err := mangaCollection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = cursor.All(context.TODO(), &mangas); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(mangas)

// 	return mangas, nil
// }
