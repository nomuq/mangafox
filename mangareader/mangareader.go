package mangareader

import (
	"context"
	"errors"
	"fmt"
	"mangafox/mal"
	"mangafox/model"
	"mangafox/store"
	"strconv"
	"time"

	"github.com/manga-community/anilist"
	"github.com/manga-community/mangareader"
	"github.com/sirupsen/logrus"
)

type SyncType int

const (
	All SyncType = iota
	Latest
)

func Sync(syncType SyncType, database string) error {
	ctx := context.Background()
	str, err := store.New(ctx, database)
	if err != nil {
		// logrus.Panic(err)
		return err
	}
	defer str.Client.Disconnect(ctx)

	mr := new(mangareader.Mangareader)

	if syncType == Latest {
		chapters, err := mr.Latest()
		if err != nil {
			return err
		}
		for _, chapter := range chapters {
			// IndexSingleChapter(str, &wg, mr, chapter)
			err := IndexChapter(str, mr, chapter)
			if err != nil {
				logrus.Errorln(err)
			}
		}

	} else if syncType == All {
		for _, link := range Links() {
			chapters, err := mr.RetrieveIssueLinks("https://www.mangareader.net/"+link, false, false)
			if err != nil {
				logrus.Errorln(err)
			}
			for _, chapter := range chapters {
				IndexChapter(str, mr, chapter)
			}

		}
	}

	str.Client.Disconnect(ctx)

	return nil
}

func IndexChapter(str *store.Store, mr *mangareader.Mangareader, chapter string) error {
	if mr.IsSingleIssue(chapter) {
		slug, chapterNumber := mr.GetInfo(chapter)

		result, err := str.GetMangaByMangareaderID(slug)
		if err != nil {
			mal, err := mal.FindFromMAL(slug)
			if err != nil {
				return err
			}

			if mal.MalID != 0 {
				anilistResult, err := anilist.GetByMAL(strconv.FormatInt(mal.MalID, 10))
				if err != nil {
					return err
				}

				var Tags []string
				if anilistResult == nil {
					return fmt.Errorf("Cant Find ON Anilist %s %s", mal.MalID, mal.Title)
				}

				for _, tag := range anilistResult.Tags {
					if tag.Name != nil {
						Tags = append(Tags, *tag.Name)
					}
				}

				MALID := strconv.FormatInt(mal.MalID, 10)
				AnilistID := strconv.FormatInt(anilistResult.ID, 10)

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
					// Chapters: emptyChepters,
				}

				_, err = str.CreateManga(record)
				if err != nil {
					return err
				}

				manga, err := str.GetMangaByMangareaderID(slug)
				if err != nil {
					return err
				}
				str.CreateMangareaderChapter(chapterNumber, manga)
			} else {
				str.CreateMangareaderMapping(slug)
			}

		} else {
			if result.Links.Mangareader == nil {
				_, err := str.UpdateMangareaderID(result, slug)
				if err != nil {
					return err
				}
			}
			str.CreateMangareaderChapter(chapterNumber, result)
		}

	} else {
		return errors.New("not a single chepter url")
	}
	return nil
}
