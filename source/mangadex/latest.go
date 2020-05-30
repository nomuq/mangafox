package mangadex

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type FeedResponse struct {
	Channel FeedChannel `xml:"channel"`
}

type FeedChannel struct {
	Items []FeedItem `xml:"item"`
}

type FeedItem struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	MangaLink string `xml:"mangaLink"`
}

func (mangadex Mangadex) Latest(token string) ([]FeedItem, error) {
	res, err := http.Get("https://mangadex.org/rss/" + token)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	var rss FeedResponse
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, err
	}
	return rss.Channel.Items, nil
}
