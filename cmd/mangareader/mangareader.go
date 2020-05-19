// package mangareader
package main

import (
	"context"
	"fmt"
	"mangafox/model"
	"mangafox/store"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/manga-community/anilist"
	"github.com/manga-community/mangareader"
	"github.com/nokusukun/jikan2go/manga"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Minute)
	str, err := store.New(ctx, "mongodb://localhost:27017")
	if err != nil {
		logrus.Panic(err)
	}
	defer str.Client.Disconnect(ctx)

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

					for _, tag := range anilistResult.Tags {
						Tags = append(Tags, *tag.Name)
					}

					// startDate := time.Date(*anilistResult.StartDate.Year, time.Month(*anilistResult.StartDate.Month), *anilistResult.StartDate.Day, 0, 0, 0, 0, time.UTC)
					// endDate := time.Date(*anilistResult.EndDate.Year, time.Month(*anilistResult.EndDate.Month), *anilistResult.EndDate.Day, 0, 0, 0, 0, time.UTC)

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
					insertResult, err := str.CreateManga(record)
					if err != nil {
						logrus.Error(err)
						return
					}
					fmt.Println(insertResult.InsertedID)
				} else {
					fmt.Println("need to index", slug)
				}

			} else {
				fmt.Println(slug, chapterNumber, result.ID)
				IndexMangareaderChepter(str, chapterNumber, result)
			}

		}
	}

	str.Client.Disconnect(ctx)
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

func IndexMangareaderChepter(str *store.Store, chepter string, record model.Manga) {

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

}
