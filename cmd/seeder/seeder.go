package main

import (
	"mangafox/store"

	"github.com/sirupsen/logrus"
)

func main() {

	store := store.Store{
		URL:    "mongodb://localhost:27017",
		DBName: "mangafox",
	}

	err := store.Connect()
	if err != nil {
		logrus.Panicln(err)
	}

	err = store.Ping()
	if err != nil {
		logrus.Panicln(err)
	}

	res, err := store.CreateMangaIndexes()
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infoln(res)
	res, err = store.CreateChapterIndexes()
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infoln(res)
	res, err = store.CreateMappingIndexes()
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infoln(res)
}
