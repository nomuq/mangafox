package mangatown

import (
	"context"
	"fmt"
	"mangafox/mal"
	"mangafox/model"
	"mangafox/store"
	"strconv"
	"strings"
	"time"

	"github.com/manga-community/anilist"
	"github.com/manga-community/mangatown"
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

	mt := new(mangatown.Mangatown)

	if syncType == Latest {
		chapters, err := mt.Latest()
		if err != nil {
			return err
		}
		for _, chapter := range chapters {
			err = IndexChapter(str, mt, chapter)
			if err != nil {
				logrus.Errorln(err)
			}
		}

	} else if syncType == All {

	}

	str.Client.Disconnect(ctx)

	return nil
}

func IndexChapter(str *store.Store, mt *mangatown.Mangatown, chapter string) error {
	mtTitle, issueNumber := mt.GetInfo(chapter)

	if strings.Contains(issueNumber, "c") {
		issueNumber = issueNumber[1:]
	}

	_, err := strconv.ParseFloat(issueNumber, 64)
	if err != nil {
		return err
	}

	malSlug := strings.ReplaceAll(mtTitle, "_", "-")

	result, err := str.GetMangaByMangatownID(mtTitle)
	if err != nil {
		mal, err := mal.FindFromMAL(malSlug)
		if err != nil {
			logrus.Errorln("MAL ERROR", mtTitle)
			return err
		}

		if mal.MalID != 0 {
			anilistResult, err := anilist.GetByMAL(strconv.FormatInt(mal.MalID, 10))
			if err != nil {
				str.CreateMangatownMapping(mtTitle)
				// logrus.Errorln("Cant Find ON Anilist", mal.MalID, mal.Title, err)
				return nil
			}

			var Tags []string
			if anilistResult == nil {
				str.CreateMangatownMapping(mtTitle)
				return nil
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
					MAL:       &MALID,
					Mangatown: &mtTitle,
					Anilist:   &AnilistID,
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
			}

			_, err = str.CreateManga(record)
			if err != nil {
				return err
			}

			manga, err := str.GetMangaByMangatownID(mtTitle)
			if err != nil {
				return err
			}
			str.CreateMangatownChapter(mt, issueNumber, manga, chapter)

		} else {
			str.CreateMangatownMapping(mtTitle)
		}

	} else {
		if result.Links.Mangatown == nil {
			_, err := str.UpdateMangatownID(result, mtTitle)
			if err != nil {
				return err
			}
		}
		str.CreateMangatownChapter(mt, issueNumber, result, chapter)
	}

	return nil
}

func SyncManga(manga string, database string) error {
	ctx := context.Background()
	str, err := store.New(ctx, database)
	if err != nil {
		// logrus.Panic(err)
		return err
	}
	defer str.Client.Disconnect(ctx)

	mt := new(mangatown.Mangatown)
	issues, err := mt.RetrieveIssueLinks("https://www.mangatown.com/manga/naruto/", false, false)

	if err != nil {
		return err
	}
	for _, chapter := range issues {
		fmt.Println(chapter)
		// err = IndexChapter(str, mt, chapter)
		// if err != nil {
		// 	logrus.Errorln(err)
		// }
	}

	str.Client.Disconnect(ctx)

	return nil
}
