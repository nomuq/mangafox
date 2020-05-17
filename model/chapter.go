package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chapter struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt string             `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt string             `json:"updatedAt" bson:"updatedAt,omitempty"`
	Title     string             `json:"title" bson:"title,omitempty"`
	Number    int32              `json:"number" bson:"number,omitempty"`
	Links     []string           `json:"links" bson:"links,omitempty"`
	Source    string             `json:"source" bson:"source,omitempty"`
	Language  string             `json:"language" bson:"language,omitempty"`
}
