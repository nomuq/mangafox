package main

import (
	"mangafox/mangareader"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	// CPU profiling by default
	// p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	var database string

	app := cli.NewApp()

	app.Name = "Mangareader Indexer"
	app.Usage = "cheptar indexer bot for mangareader"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "database",
			Value:       "mongodb://localhost:27017",
			Usage:       "mongo db url",
			Destination: &database,
		},
	}

	app.Action = func(c *cli.Context) error {
		logrus.Infoln("Indexing Latest Chapters")
		err := mangareader.Sync(mangareader.Latest, database)
		return err
	}

	app.Commands = []*cli.Command{
		{
			Name:    "latest",
			Aliases: []string{"l"},
			Usage:   "index latest mangareader cheptars",
			Action: func(c *cli.Context) error {
				logrus.Infoln("Indexing Latest Chapters")

				err := mangareader.Sync(mangareader.Latest, database)
				return err
			},
		},
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "index all mangareader cheptars",
			Action: func(c *cli.Context) error {
				logrus.Infoln("Indexing All Chapters")

				err := mangareader.Sync(mangareader.All, database)
				return err
			},
		},
		{
			Name:    "single",
			Aliases: []string{"s"},
			Usage:   "index single mangareader cheptars",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "manga",
					Value:    "naruto",
					Usage:    "manga title",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				logrus.Infoln("Indexing Single Chapters")

				err := mangareader.SyncManga("naruto", database)
				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}

	// p.Stop()

}
