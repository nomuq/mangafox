package main

import (
	"mangafox/source/mangadex"
	"mangafox/tasks"
	"path"
	"time"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func main() {
	options := asynq.RedisClientOpt{Addr: "localhost:6379"}
	queue := asynq.NewClient(options)

	err := Latest("2ZevhabKgkstB6DPzQpMcdSRnxwf78uC", queue)
	if err != nil {
		logrus.Panicln(err)
	}

	//md := mangadex.Initilize()
	//m, err := md.GetInfo("5")
	//if err != nil {
	//	logrus.Panicln(err)
	//}
	//for _, item := range m.Chapters {
	//	payload := map[string]interface{}{"manga_id": "5", "chapter_id": item.ID}
	//	task := asynq.NewTask(string(tasks.IndexMangadexChapter), payload)
	//	queue.Enqueue(task, asynq.Unique(time.Hour), asynq.MaxRetry(2))
	//}
}

func Latest(token string, queue *asynq.Client) error {
	md := mangadex.Initilize()
	items, err := md.Latest(token)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	for _, item := range items {
		chapter := path.Base(item.Link)
		manga := path.Base(item.MangaLink)
		err = IndexChapter(manga, chapter, queue)
		if err != nil {
			logrus.Errorln(err)
		}
	}

	return nil
}

func IndexChapter(manga string, chapter string, queue *asynq.Client) error {
	payload := map[string]interface{}{"manga_id": manga, "chapter_id": chapter}
	task := asynq.NewTask(string(tasks.IndexMangadexChapter), payload)
	err := queue.Enqueue(task, asynq.Unique(time.Hour), asynq.MaxRetry(2))
	return err
}
