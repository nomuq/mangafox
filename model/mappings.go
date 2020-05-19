package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mapping struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Slug      string             `json:"slug" bson:"slug"`
	Indexed   bool               `json:"indexed" bson:"indexed"`
	Source    string             `json:"source" bson:"source"`
	Language  string             `json:"language" bson:"language"`
}
