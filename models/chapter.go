package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Chapter struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Title     *string            `json:"title" bson:"title"`
	Number    float64            `json:"number" bson:"number"`
	Links     []string           `json:"links" bson:"links"`
	Source    string             `json:"source" bson:"source"`
	Language  string             `json:"language" bson:"language"`
	Manga     primitive.ObjectID `json:"manga" bson:"manga"`
}
