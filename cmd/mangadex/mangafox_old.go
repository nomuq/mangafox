// package main

// import (
// 	"os"
// 	"os/signal"
// 	"path"

// 	"github.com/gocraft/work"
// 	"github.com/gomodule/redigo/redis"
// 	"github.com/manga-community/mangadex"
// 	"github.com/sirupsen/logrus"
// )

// type Context struct {
// 	customerID int64
// 	mangadex   *mangadex.Mangadex
// 	enqueuer   *work.Enqueuer
// }

// const (
// 	Port = ":8080"
// )

// func main() {

// 	redisPool := &redis.Pool{
// 		MaxActive: 5,
// 		MaxIdle:   5,
// 		Wait:      true,
// 		Dial: func() (redis.Conn, error) {
// 			return redis.Dial("tcp", ":6379")
// 		},
// 	}

// 	md := mangadex.Initilize()

// 	var enqueuer = work.NewEnqueuer("mangafox_indexer", redisPool)

// 	ctx := Context{
// 		mangadex: md,
// 		enqueuer: enqueuer,
// 	}

// 	pool := work.NewWorkerPool(ctx, 10, "mangafox_indexer_worker", redisPool)
// 	pool.PeriodicallyEnqueue("0 * * * * *", "index_latest")
// 	pool.Job("index_latest", ctx.Latest)
// 	pool.Job("index_chapter", ctx.IndexChapter)
// 	pool.Start()

// 	// Wait for a signal to quit:
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, os.Interrupt, os.Kill)
// 	<-signalChan

// 	// Stop the pool
// 	pool.Stop()

// }

// func (ctx *Context) Latest(job *work.Job) error {
// 	logrus.Infoln("Indexing Latest")
// 	links, err := ctx.mangadex.Latest("2ZevhabKgkstB6DPzQpMcdSRnxwf78uC")
// 	if err != nil {
// 		logrus.Errorln(err)
// 		return err
// 	}
// 	for _, element := range links {
// 		id := path.Base(element)
// 		job, err := ctx.enqueuer.Enqueue("index_chapter", work.Q{"object_id_": id, "chapter": id})
// 		if err != nil {
// 			logrus.Errorln(err)
// 			return err
// 		}
// 		logrus.Infoln("Enqueued ", job.ID)
// 	}

// 	return nil
// }

// func (ctx *Context) IndexChapter(job *work.Job) error {
// 	logrus.Infoln("Indexing chapter ")
// 	chapter := job.ArgString("chapter")
// 	if err := job.ArgError(); err != nil {
// 		logrus.Errorln(err)
// 		return err
// 	}

// 	logrus.Infoln("Indexing chapter ", chapter)
// 	return nil
// }
