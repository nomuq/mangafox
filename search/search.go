package search

import (
	"mangafox/models"

	"github.com/meilisearch/meilisearch-go"
	"github.com/sirupsen/logrus"
)

const MangaIndex = "manga"

type Search struct {
	URL    string
	client *meilisearch.Client
}

func (search *Search) Initilize() {
	client := meilisearch.NewClient(meilisearch.Config{
		Host: search.URL,
	})
	search.client = client
}

func (search *Search) CreateIndexes() error {
	indexes, err := search.client.Indexes().List()
	if err != nil {
		return err
	}
	if len(indexes) != 0 {
		return nil
	}

	result, err := search.client.Indexes().Create(meilisearch.CreateIndexRequest{
		UID: MangaIndex,
	})
	if err != nil {
		return err
	}
	logrus.Infoln("Created Search Index For Manga", result.Name)

	return nil
}

func (search *Search) IndexAllManga(documents []models.Manga) (*meilisearch.AsyncUpdateID, error) {
	result, err := search.client.Documents(MangaIndex).AddOrUpdate(documents)
	return result, err
}
