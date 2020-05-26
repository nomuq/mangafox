package main

import (
	"os"
	"path"

	"github.com/manga-community/mangadex"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	var database string
	var token string
	var chapter string

	app := cli.NewApp()

	app.Name = "Mangadex Indexer"
	app.Usage = "cheptar indexer bot for mangadex"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "mongo, m",
			Value:       "mongodb://localhost:27017",
			Usage:       "mongo db uri",
			EnvVars:     []string{"MONGO_URI"},
			Destination: &database,
		},
		&cli.StringFlag{
			Name:        "token, t",
			Value:       "2ZevhabKgkstB6DPzQpMcdSRnxwf78uC",
			Usage:       "rss token",
			EnvVars:     []string{"RSS_TOKEN"},
			Destination: &token,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "latest",
			Aliases: []string{"l"},
			Usage:   "index latest mangadex cheptars",
			Action: func(c *cli.Context) error {

				md := mangadex.Initilize()
				err := Latest(token, md)
				return err
			},
		},
		{
			Name:    "chapter",
			Aliases: []string{"c"},
			Usage:   "index mangadex chapter",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "id",
					Usage:       "chapter id",
					Required:    true,
					Destination: &chapter,
				},
			},
			Action: func(c *cli.Context) error {
				md := mangadex.Initilize()
				err := IndexChapter(chapter, md)
				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func Latest(token string, mangadex *mangadex.Mangadex) error {
	logrus.Infoln("Indexing Latest Chapters")
	links, err := mangadex.Latest("2ZevhabKgkstB6DPzQpMcdSRnxwf78uC")
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	for _, element := range links {
		id := path.Base(element)
		logrus.Infoln("Enqueued ", id)
	}

	return nil
}

func IndexChapter(chapter string, mangadex *mangadex.Mangadex) error {
	logrus.Infoln("Indexing ", chapter)

	return nil
}
