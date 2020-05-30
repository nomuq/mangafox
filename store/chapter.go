package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mangafox/models"
	"time"
)

const ChapterCollection = "chapter"

func (store Store) CreateChapter(chapter models.Chapter) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	result, err := store.db.Collection(ChapterCollection).InsertOne(ctx, chapter)
	if err != nil {
		return primitive.NewObjectID(), err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return objectID, err
	}

	return primitive.NewObjectID(), err
}
