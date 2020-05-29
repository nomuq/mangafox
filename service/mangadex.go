package service

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

func (service *Service) IndexMangadexChapter(ctx context.Context, t *asynq.Task) error {

	mangaID, err := t.Payload.GetString("manga_id")
	if err != nil {
		return err
	}

	if mangaID == "" {
		return fmt.Errorf("Empty mangadex manga id")
	}

	chapterID, err := t.Payload.GetString("chapter_id")
	if err != nil {
		return err
	}

	if chapterID == "" {
		return fmt.Errorf("Empty mangadex chapter id")
	}

	// service.Store.Context = ctx
	// fmt.Println(service.Store.Context)

	manga, err := service.Store.GetMangaByMangadexID(chapterID)
	if err != nil {
		// Create Manga and Index Chapter

	} else {
		// INDEX Chepter
	}
	fmt.Println(manga.ID)

	return nil
}
