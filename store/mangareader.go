package store

import (
	"fmt"
	"mangafox/model"
	"strconv"
	"time"

	"github.com/manga-community/mangareader"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *Store) GetMangaByMangareaderID(slug string) (model.Manga, error) {
	var result model.Manga
	filter := bson.D{primitive.E{Key: "links.mangareader", Value: slug}}
	err := store.MangaCollection().FindOne(store.ctx, filter).Decode(&result)
	return result, err
}

func (store *Store) CreateMangareaderMapping(slug string) (*mongo.UpdateResult, error) {
	record := model.Mapping{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Language:  "en",
		Source:    "mangareader",
		Slug:      slug,
	}
	result, err := store.CreateMapping(record)
	return result, err
}

func (store *Store) CreateMangareaderChapter(issueNumber string, manga model.Manga) (*mongo.InsertOneResult, error) {
	URL := fmt.Sprintf("https://www.mangareader.net/%s/%s/", *manga.Links.Mangareader, issueNumber)
	SOURCE := "www.mangareader.net"

	mr := mangareader.Mangareader{}

	comic := new(mangareader.Comic)
	comic.Name = *manga.Links.Mangareader
	comic.IssueNumber = issueNumber
	comic.URLSource = URL
	comic.Source = SOURCE

	links, err := mr.RetrieveImageLinks(comic)
	if err != nil {
		return nil, err
	}

	number, err := strconv.ParseInt(issueNumber, 10, 64)
	if err != nil {
		return nil, err
	}

	chapter := model.Chapter{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Language:  "en",
		Number:    number,
		Source:    "mangareader",
		Links:     links,
		Manga:     manga.ID,
	}

	result, err := store.CreateChapter(manga, chapter)
	logrus.Infoln("Indexed", manga.Title, number)
	return result, err
}
