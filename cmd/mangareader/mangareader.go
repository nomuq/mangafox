// package mangareader
package main

import (
	"context"
	"fmt"
	"mangafox/model"
	"mangafox/store"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gosimple/slug"
	"github.com/manga-community/anilist"
	"github.com/manga-community/mangareader"
	"github.com/nokusukun/jikan2go/manga"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Minute)
	ctx := context.Background()
	str, err := store.New(ctx, "mongodb://localhost:27017")
	if err != nil {
		logrus.Panic(err)
	}
	defer str.Client.Disconnect(ctx)

	SyncLatestChapters(str)

	str.Client.Disconnect(ctx)
}

func SyncLatestChapters(str *store.Store) {
	var wg sync.WaitGroup

	mr := new(mangareader.Mangareader)
	chapters, err := mr.Latest()
	if err != nil {
		logrus.Panic(err)
	}
	for _, chapter := range chapters {
		if mr.IsSingleIssue(chapter) {
			slug, chapterNumber := mr.GetInfo(chapter)

			result, err := str.GetMangaByMangareaderID(slug)
			if err != nil {
				mal, err := FindFromMAL(slug)
				if err != nil {
					logrus.Error(err)
				}

				if mal.MalID != 0 {
					anilistResult, _ := anilist.GetByMAL(strconv.FormatInt(mal.MalID, 10))
					var Tags []string

					if anilistResult == nil {
						logrus.Warnln("Cant Find ON Anilist", mal.MalID, mal.Title)
						return
					}

					for _, tag := range anilistResult.Tags {
						if tag.Name != nil {
							Tags = append(Tags, *tag.Name)
						}
					}

					MALID := strconv.FormatInt(mal.MalID, 10)
					AnilistID := strconv.FormatInt(anilistResult.ID, 10)
					emptyChepters := make([]primitive.ObjectID, 0)

					record := model.Manga{
						Title: mal.Title,
						Type:  string(mal.Type),

						Description:  mal.Synopsis,
						IsPublishing: mal.Publishing,
						Links: model.Links{
							MAL:         &MALID,
							Mangareader: &slug,
							Anilist:     &AnilistID,
						},
						Genres:   anilistResult.Genres,
						Tags:     Tags,
						Synonyms: anilistResult.Synonyms,
						Cover: model.Cover{
							Color:      anilistResult.CoverImage.Color,
							ExtraLarge: anilistResult.CoverImage.ExtraLarge,
							Large:      anilistResult.CoverImage.Large,
							Medium:     anilistResult.CoverImage.Medium,
						},
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Banner:    anilistResult.BannerImage,
						StartDate: model.Date{
							Day:   anilistResult.StartDate.Day,
							Month: anilistResult.StartDate.Month,
							Year:  anilistResult.StartDate.Year,
						},
						EndDate: model.Date{
							Day:   anilistResult.EndDate.Day,
							Month: anilistResult.EndDate.Month,
							Year:  anilistResult.EndDate.Year,
						},
						Chapters: emptyChepters,
					}

					_, err := str.CreateManga(record)
					if err != nil {
						logrus.Error(err)
						return
					}

					result, err := str.GetMangaByMangareaderID(slug)
					if err != nil {
						logrus.Error(err)
						return
					}
					wg.Add(1)
					go IndexMangareaderChepter(str, chapterNumber, result, &wg)
				} else {
					wg.Add(1)
					go CreateMapping(str, slug, &wg)
				}

			} else {
				wg.Add(1)
				go IndexMangareaderChepter(str, chapterNumber, result, &wg)
			}

		}
	}

	wg.Wait()
}

func FindFromMAL(title string) (manga.Result, error) {
	result, err := manga.Search(manga.Query{Q: title})

	if err != nil {
		return manga.Result{}, err
	}

	for _, r := range result.Results {
		if strings.ToUpper(slug.Make(r.Title)) == strings.ToUpper(title) {
			return r, err
		}
	}
	return manga.Result{}, err
}

func IndexMangareaderChepter(str *store.Store, chepter string, record model.Manga, wg *sync.WaitGroup) {
	defer wg.Done()
	// logrus.Infoln(record.Title, chepter)

	URL := fmt.Sprintf("https://www.mangareader.net/%s/%s/", *record.Links.Mangareader, chepter)
	SOURCE := "www.mangareader.net"

	mr := mangareader.Mangareader{}

	comic := new(mangareader.Comic)
	comic.Name = *record.Links.Mangareader
	comic.IssueNumber = chepter
	comic.URLSource = URL
	comic.Source = SOURCE

	links, err := mr.RetrieveImageLinks(comic)
	if err != nil {
		logrus.Error(err)
		return
	}
	if number, err := strconv.ParseInt(chepter, 10, 64); err == nil {
		chapterRecord := model.Chapter{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Language:  "en",
			Number:    number,
			Source:    "mangareader",
			Links:     links,
			Manga:     record.ID,
		}

		str.CreateChapter(record, chapterRecord)
	}
	logrus.Infoln("INDEXED", record.Title, chepter)

}

func CreateMapping(str *store.Store, slug string, wg *sync.WaitGroup) {
	defer wg.Done()
	record := model.Mapping{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Language:  "en",
		Source:    "mangareader",
		Slug:      slug,
	}
	_, err := str.CreateMapping(record)
	if err != nil {
		logrus.Error(err)
	}
}
