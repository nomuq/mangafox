package store

import (
	"fmt"
	"mangafox/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *Store) MangaCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("manga")
}

func (store *Store) ChapterCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("chapter")
}

func (store *Store) GetMangaByMangareaderID(slug string) (model.Manga, error) {
	var result model.Manga
	filter := bson.D{primitive.E{Key: "links.mangareader", Value: slug}}
	err := store.MangaCollection().FindOne(store.ctx, filter).Decode(&result)
	return result, err
}

func (store *Store) CreateManga(manga model.Manga) (*mongo.InsertOneResult, error) {
	result, err := store.MangaCollection().InsertOne(store.ctx, manga)
	return result, err
}

func (store *Store) CreateChapter(manga model.Manga, chapter model.Chapter) (*mongo.InsertOneResult, error) {
	// mangaCollection := str.MangaCollection()
	// chapterCollection := str.ChapterCollection()

	// opts := options.Update().SetUpsert(true)
	// filter := bson.D{{"manga", manga.ID}}

	// filter := bson.D{primitive.E{Key: "manga", Value: slug}}

	// store.ChapterCollection().UpdateOne()

	result, err := store.ChapterCollection().InsertOne(store.ctx, chapter)

	return result, err
}

func (store *Store) GetAllManga() ([]model.Manga, error) {

	database := store.Client.Database("mangafox")
	mangaCollection := database.Collection("manga")
	// chapetrsCollection := database.Collection("chapters")

	var mangas []model.Manga
	cursor, err := mangaCollection.Find(store.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(store.ctx, &mangas); err != nil {
		panic(err)
	}
	fmt.Println(mangas)

	return mangas, nil
}
