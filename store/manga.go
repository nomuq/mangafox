package store

import (
	"context"
	"fmt"
	"mangafox/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (store *Store) MangaIndexes() {
	cursor, err := store.MangaCollection().Indexes().List(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	for cursor.Next(context.TODO()) {
		index := bson.D{}
		cursor.Decode(&index)
		fmt.Println(fmt.Sprintf("index found %v", index))
	}
}

func (store *Store) MangaCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("manga")
}

func (store *Store) ChapterCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("chapter")
}

func (store *Store) MappingsCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("mapping")
}

func (store *Store) GetMangaByMangareaderID(slug string) (model.Manga, error) {
	var result model.Manga
	filter := bson.D{primitive.E{Key: "links.mangareader", Value: slug}}
	err := store.MangaCollection().FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (store *Store) CreateManga(manga model.Manga) (*mongo.InsertOneResult, error) {
	result, err := store.MangaCollection().InsertOne(context.TODO(), manga)
	return result, err
}

func (store *Store) CreateChapter(manga model.Manga, chapter model.Chapter) (*mongo.InsertOneResult, error) {
	filter := bson.M{
		"$and": []bson.M{
			bson.M{"language": "en"},
			bson.M{"manga": manga.ID},
			bson.M{"source": chapter.Source},
			bson.M{"number": chapter.Number},
		},
	}

	// store.ChapterCollection().UpdateOne()
	var r model.Chapter
	err := store.ChapterCollection().FindOne(context.TODO(), filter).Decode(&r)

	if err != nil {
		result, err := store.ChapterCollection().InsertOne(context.TODO(), chapter)
		return result, err
	}

	return nil, err
}

func (store *Store) GetAllManga() ([]model.Manga, error) {

	database := store.Client.Database("mangafox")
	mangaCollection := database.Collection("manga")
	// chapetrsCollection := database.Collection("chapters")

	var mangas []model.Manga
	cursor, err := mangaCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &mangas); err != nil {
		panic(err)
	}
	fmt.Println(mangas)

	return mangas, nil
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
