package main

import (
	"mangafox/mangatown"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	// CPU profiling by default
	// p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	var database string

	app := cli.NewApp()

	app.Name = "Mangatown Indexer"
	app.Usage = "cheptar indexer bot for mangatown"

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
		err := mangatown.Sync(mangatown.Latest, database)
		return err
	}

	app.Commands = []*cli.Command{
		{
			Name:    "latest",
			Aliases: []string{"l"},
			Usage:   "index latest mangareader cheptars",
			Action: func(c *cli.Context) error {
				logrus.Infoln("Indexing Latest Chapters")

				err := mangatown.Sync(mangatown.Latest, database)
				return err
			},
		},
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "index all mangareader cheptars",
			Action: func(c *cli.Context) error {
				logrus.Infoln("Indexing All Chapters")

				err := mangatown.Sync(mangatown.All, database)
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
