package main

import (
	"mangafox/tasks"
	"os"
	"path"
	"time"

	"github.com/hibiken/asynq"
	"github.com/manga-community/mangadex"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	var redis string
	var token string
	// var chapter string

	app := cli.NewApp()

	app.Name = "Mangadex Indexer"
	app.Usage = "cheptar indexer bot for mangadex"

	app.Flags = []cli.Flag{
		// &cli.StringFlag{
		// 	Name:        "mongo, m",
		// 	Value:       "mongodb://localhost:27017",
		// 	Usage:       "mongo db uri",
		// 	EnvVars:     []string{"MONGO_URI"},
		// 	Destination: &database,
		// },
		&cli.StringFlag{
			Name:        "redis, r",
			Value:       "localhost:6379",
			Usage:       "redis url",
			EnvVars:     []string{"REDIS_URL"},
			Destination: &redis,
		},
		&cli.StringFlag{
			Name:        "token, t",
			Value:       "2ZevhabKgkstB6DPzQpMcdSRnxwf78uC",
			Usage:       "rss token",
			EnvVars:     []string{"RSS_TOKEN"},
			Destination: &token,
		},
	}

	app.Action = func(c *cli.Context) error {
		err := Latest(token, redis)
		return err
	}

	app.Commands = []*cli.Command{
		{
			Name:    "latest",
			Aliases: []string{"l"},
			Usage:   "index latest mangadex cheptars",
			Action: func(c *cli.Context) error {
				err := Latest(token, redis)
				return err
			},
		},
		// {
		// 	Name:    "chapter",
		// 	Aliases: []string{"c"},
		// 	Usage:   "index mangadex chapter",
		// 	Flags: []cli.Flag{
		// 		&cli.StringFlag{
		// 			Name:        "id",
		// 			Usage:       "chapter id",
		// 			Required:    true,
		// 			Destination: &chapter,
		// 		},
		// 	},
		// 	Action: func(c *cli.Context) error {
		// 		err := IndexSingleChapter(chapter)
		// 		return err
		// 	},
		// },
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func Latest(token string, redis string) error {

	options := asynq.RedisClientOpt{Addr: redis}
	queue := asynq.NewClient(options)

	md := mangadex.Initilize()

	items, err := md.Latest(token)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	for _, item := range items {
		chapter := path.Base(item.Link)
		manga := path.Base(item.MangaLink)
		err = IndexChapter(manga, chapter, md, queue)
		if err != nil {
			logrus.Errorln(err)
		}
	}

	return nil
}

func IndexChapter(manga string, chapter string, mangadex *mangadex.Mangadex, queue *asynq.Client) error {
	payload := map[string]interface{}{"manga_id": manga, "chapter_id": chapter}
	task := asynq.NewTask(string(tasks.IndexMangadexChapter), payload)
	err := queue.Enqueue(task, asynq.Unique(time.Hour))
	return err
	return nil
}
