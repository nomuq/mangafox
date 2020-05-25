package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Manga struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
	Title          string             `json:"title" bson:"title"`
	AlternateTitle AlternateTitle     `json:"alternateTitle" bson:"alternateTitle"`
	Description    string             `json:"description" bson:"description"`
	Genres         []string           `json:"genres" bson:"genres"`
	Tags           []string           `json:"tags" bson:"tags"`
	Synonyms       []string           `json:"synonyms" bson:"synonyms"`
	Type           string             `json:"type" bson:"type"`
	Banner         *string            `json:"banner" bson:"banner"`
	IsPublishing   bool               `json:"isPublishing" bson:"isPublishing"`
	StartDate      Date               `json:"startDate" bson:"startDate"`
	EndDate        Date               `json:"endDate" bson:"endDate"`
	Links          Links              `json:"links" bson:"links"`
	Cover          Cover              `json:"cover" bson:"cover"`
	Chapters       []Chapter          `json:"chapters" bson:"-"`
	Country        *string            `json:"country" bson:"country"`
}

type Links struct {
	Anilist     *string `json:"anilist"  bson:"anilist"`
	MAL         *string `json:"mal" bson:"mal"`
	Mangadex    *string `json:"mangadex" bson:"mangadex"`
	Mangareader *string `json:"mangareader" bson:"mangareader"`
	Mangatown   *string `json:"mangatown" bson:"mangatown"`
}

type Cover struct {
	ExtraLarge *string `json:"extraLarge"`
	Large      *string `json:"large"`
	Medium     *string `json:"medium"`
	Color      *string `json:"color"`
}

type Date struct {
	Year  *int `json:"year"`
	Month *int `json:"month"`
	Day   *int `json:"day"`
}

type AlternateTitle struct {
	Romaji  *string `json:"romaji"`
	English *string `json:"english"`
	Native  *string `json:"native"`
}
