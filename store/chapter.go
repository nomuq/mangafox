package store

import (
	"context"
	"mangafox/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *Store) ChapterCollection() *mongo.Collection {
	return store.Client.Database("mangafox").Collection("chapter")
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
