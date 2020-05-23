package mangatown

import (
	"context"
	"fmt"
	"mangafox/store"

	"github.com/manga-community/mangatown"
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
			fmt.Println(chapter)
			IndexChapter(str, mt, chapter)
		}

	} else if syncType == All {

	}

	str.Client.Disconnect(ctx)

	return nil
}

func IndexChapter(str *store.Store, mt *mangatown.Mangatown, chapter string) error {
	slug, chapterNumber := mt.GetInfo(chapter)
	fmt.Println(slug, chapterNumber)

	return nil
}
