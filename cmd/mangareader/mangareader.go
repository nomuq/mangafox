// package mangareader
package main

import (
	"context"
	"fmt"
	"mangafox/store"
	"time"

	"github.com/manga-community/mangareader"
	"github.com/sirupsen/logrus"
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

			} else {
				fmt.Println(slug, chapterNumber, result.ID)
			}

		}
	}

	str.Client.Disconnect(ctx)
}
