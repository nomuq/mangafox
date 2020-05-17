package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manga struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	CreatedAt    string               `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt    string               `json:"updatedAt" bson:"updatedAt,omitempty"`
	Title        string               `json:"title" bson:"title,omitempty"`
	Description  string               `json:"description" bson:"description,omitempty`
	Genres       []string             `json:"genres" bson:"genres,omitempty`
	Tags         []string             `json:"tags" bson:"tags,omitempty`
	Synonyms     []string             `json:"synonyms" bson:"synonyms,omitempty`
	Type         string               `json:"type" bson:"type,omitempty`
	Banner       string               `json:"banner" bson:"banner,omitempty`
	IsPublishing bool                 `json:"isPublishing" bson:"isPublishing,omitempty"`
	StartDate    string               `json:"startDate" bson:"startDate,omitempty"`
	EndDate      string               `json:"endDate" bson:"endDate,omitempty"`
	Chapters     []primitive.ObjectID `json:"chapters" bson:"chapters,omitempty"`
}
