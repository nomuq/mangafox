package worker

import (
	"context"
	"fmt"
	"mangafox/models"
	"mangafox/source/anilist"
	"mangafox/source/mangadex"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hibiken/asynq"
)

func (worker Worker) IndexMangadexChapter(ctx context.Context, t *asynq.Task) error {

	mangaID, err := t.Payload.GetString("manga_id")
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	if mangaID == "" {
		logrus.Errorln(fmt.Errorf("empty mangadex manga id"))
		return fmt.Errorf("empty mangadex manga id")
	}

	chapterID, err := t.Payload.GetString("chapter_id")
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	if chapterID == "" {
		logrus.Errorln(fmt.Errorf("empty mangadex chapter id"))
		return fmt.Errorf("empty mangadex chapter id")
	}

	// Check if chapter is already indexed
	cacheKey := "mangadex" + ":" + mangaID + ":" + chapterID
	//isChapterIndexed, err := worker.cache.GetBool(cacheKey)
	//if isChapterIndexed {
	//	return nil
	//}

	md := mangadex.Initilize()
	manga, err := worker.store.GetMangaByMangadexID(mangaID)

	if err != nil {

		mangadexManga, err := md.GetInfo(mangaID)
		if err != nil {
			logrus.Errorln(err)
			return err
		}

		if mangadexManga.Hentai != 0 {
			return nil
		}

		var anilistResult anilist.Manga
		var Tags []string
		Description := mangadexManga.Description

		if mangadexManga.Links.Al != "" {
			anilistResult, err = anilist.GetByID(mangadexManga.Links.Al)
			if err == nil {
				Description = *anilistResult.Description
				for _, tag := range anilistResult.Tags {
					if tag.Name != nil {
						Tags = append(Tags, *tag.Name)
					}
				}
			}
		}

		IsPublishing := true
		if mangadexManga.Status == 2 {
			IsPublishing = false
		}

		record := models.Manga{
			Title:        mangadexManga.Title,
			Description:  Description,
			IsPublishing: IsPublishing,
			Links: models.Links{
				Mangadex: &mangaID,
				MAL:      &mangadexManga.Links.Mal,
				Anilist:  &mangadexManga.Links.Al,
			},

			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),

			Author: mangadexManga.Author,
			Artist: mangadexManga.Artist,

			// Anilist Data
			Tags:     Tags,
			Genres:   anilistResult.Genres,
			Synonyms: anilistResult.Synonyms,
			Cover: models.Cover{
				Color:      anilistResult.CoverImage.Color,
				ExtraLarge: anilistResult.CoverImage.ExtraLarge,
				Large:      anilistResult.CoverImage.Large,
				Medium:     anilistResult.CoverImage.Medium,
				Default:    mangadexManga.CoverURL,
			},
			Banner: anilistResult.BannerImage,
			StartDate: models.Date{
				Day:   anilistResult.StartDate.Day,
				Month: anilistResult.StartDate.Month,
				Year:  anilistResult.StartDate.Year,
			},
			EndDate: models.Date{
				Day:   anilistResult.EndDate.Day,
				Month: anilistResult.EndDate.Month,
				Year:  anilistResult.EndDate.Year,
			},
			AlternateTitle: models.AlternateTitle{
				English: anilistResult.Title.English,
				Native:  anilistResult.Title.Native,
				Romaji:  anilistResult.Title.Romaji,
			},
			Country: anilistResult.CountryOfOrigin,
		}

		//if mangadexManga.Links.Mal != "" {
		//	malID, err := strconv.ParseInt(mangadexManga.Links.Mal, 0, 64)
		//	if err == nil {
		//		malManga, err := mal.GetManga(malID)
		//		if err == nil {
		//			manga.Type = malManga.Type
		//			manga.IsPublishing = malManga.Publishing
		//		}
		//	}
		//}

		recordID, err := worker.store.CreateManga(record)
		if err != nil {
			logrus.Errorln(err)
			return err
		}

		manga.ID = recordID
	}

	if manga.ID == primitive.NilObjectID {
		logrus.Errorln(fmt.Errorf("manga object id is empty"))
		return fmt.Errorf("manga object id is empty")
	}

	chapter, err := md.RetrieveImageLinks(chapterID)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	number, err := strconv.ParseFloat(chapter.Number, 64)
	if err != nil {
		logrus.Errorln(err)
		return nil
	}

	record := models.Chapter{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Language:  chapter.Language,
		Number:    number,
		Source:    "mangadex",
		Links:     chapter.Links,
		Manga:     manga.ID,
		Title:     &chapter.Title,
	}

	result, err := worker.store.CreateChapter(record)
	if err != nil {
		logrus.Errorln(err)
		return nil
	}
	logrus.Infoln(cacheKey, result)

	//worker.cache.SetBool(cacheKey, true)

	return nil
}
