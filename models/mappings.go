package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mapping struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Manga     string             `json:"manga" bson:"manga"`
	Chapter   string             `json:"chapter" bson:"chapter"`
	Source    string             `json:"source" bson:"source"`
	Language  string             `json:"language" bson:"language"`
	Indexed   bool               `json:"indexed" bson:"indexed"`
	// Chapter  primitive.ObjectID `json:"chapter" bson:"chapter"`
}
