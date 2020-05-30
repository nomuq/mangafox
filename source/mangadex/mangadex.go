package mangadex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Mangadex struct {
	client *http.Client
}

type response struct {
	Manga   Manga              `json:"manga"`
	Chapter map[string]Chapter `json:"chapter"`
}

func Initilize() Mangadex {
	mangadex := Mangadex{
		client: http.DefaultClient,
	}
	return mangadex
}

func (mangadex Mangadex) GetInfo(id string) (Manga, error) {
	var resp response

	url := "https://mangadex.org/api/manga/" + id

	res, err := mangadex.client.Get(url)
	if err != nil {
		return resp.Manga, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return resp.Manga, fmt.Errorf("could not get %s: %s", url, res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp.Manga, err
	}

	var Chapters []Chapter

	for key, element := range resp.Chapter {
		element.ID, _ = strconv.ParseInt(key, 10, 32)
		Chapters = append(Chapters, element)
	}

	resp.Manga.Chapters = Chapters
	resp.Manga.CoverURL = "https://mangadex.org" + resp.Manga.CoverURL

	return resp.Manga, err
}

func (mangadex Mangadex) RetrieveImageLinks(id string) (Chapter, error) {
	var chapter Chapter

	url := "https://mangadex.org/api/?type=chapter&id=" + id

	res, err := mangadex.client.Get(url)
	if err != nil {
		return chapter, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return chapter, fmt.Errorf("could not get %s: %s", url, res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(data, &chapter)
	if err != nil {
		return chapter, err
	}

	var Links []string
	for _, element := range chapter.Links {
		link := chapter.Server + chapter.Hash + "/" + element
		Links = append(Links, link)
	}

	chapter.Links = Links

	return chapter, err
}
