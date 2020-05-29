package store

import (
	"mangafox/model"
	"net/url"
	"strconv"
	"time"

	"github.com/manga-community/mangatown"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (store *Store) GetMangaByMangatownID(slug string) (model.Manga, error) {
	var result model.Manga
	filter := bson.D{primitive.E{Key: "links.mangatown", Value: slug}}
	err := store.MangaCollection().FindOne(store.Context, filter).Decode(&result)
	return result, err
}

func (store *Store) UpdateMangatownID(manga model.Manga, slug string) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{Key: "_id", Value: manga.ID}}
	update := bson.D{primitive.E{Key: "$set",
		Value: bson.D{
			primitive.E{Key: "links.mangatown", Value: slug},
		},
	}}

	result, err := store.MangaCollection().UpdateOne(store.Context, filter, update, opts)
	return result, err
}

func (store *Store) CreateMangatownMapping(slug string) (*mongo.UpdateResult, error) {
	record := model.Mapping{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Language:  "en",
		Source:    "mangatown",
		Slug:      slug,
	}
	result, err := store.CreateMapping(record)
	return result, err
}

func (store *Store) CreateMangatownChapter(mt *mangatown.Mangatown, issueNumber string, manga model.Manga, link string) (*mongo.InsertOneResult, error) {

	comic := new(mangatown.Comic)
	comic.URLSource = link
	err := mt.Initialize(comic)
	if err != nil {
		return nil, err
	}
	var Links []string

	for _, item := range comic.Links {
		u, err := url.Parse(item)
		if err == nil {
			resultURL := u.Scheme + "://" + u.Host + u.Path
			Links = append(Links, resultURL)
		}
	}

	number, err := strconv.ParseFloat(issueNumber, 64)
	if err != nil {
		return nil, err
	}

	chapter := model.Chapter{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Language:  "en",
		Number:    number,
		Source:    "mangatown",
		Links:     Links,
		Manga:     manga.ID,
	}

	result, err := store.CreateChapter(manga, chapter)
	logrus.Infoln("Indexed", manga.Title, number)
	return result, err
}
