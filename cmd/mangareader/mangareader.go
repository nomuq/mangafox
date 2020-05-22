package main

import (
	"fmt"
	"html"
	"mangafox/mangareader"
	"os"

	"net/http"
	_ "net/http/pprof"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
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

				http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
					mangareader.Sync(mangareader.Latest, database)
					fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
				})

				http.ListenAndServe("localhost:8080", nil)

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

				http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
					mangareader.Sync(mangareader.All, database)
					fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
				})

				http.ListenAndServe("localhost:8080", nil)

				err := mangareader.Sync(mangareader.All, database)
				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}

}

// func IndexAll() {
// 	ctx := context.Background()
// 	str, err := store.New(ctx, "mongodb://localhost:27017")
// 	if err != nil {
// 		logrus.Panic(err)
// 	}
// 	defer str.Client.Disconnect(ctx)

// 	SyncAllChapters(str)

// 	str.Client.Disconnect(ctx)
// }

// func IndexLatest() {
// 	ctx := context.Background()
// 	str, err := store.New(ctx, "mongodb://localhost:27017")
// 	if err != nil {
// 		logrus.Panic(err)
// 	}
// 	defer str.Client.Disconnect(ctx)

// 	SyncLatestChapters(str)

// 	str.Client.Disconnect(ctx)
// }

// func SyncLatestChapters(str *store.Store) {
// 	var wg sync.WaitGroup

// 	mr := new(mangareader.Mangareader)
// 	chapters, err := mr.Latest()
// 	if err != nil {
// 		logrus.Panic(err)
// 	}
// 	for _, chapter := range chapters {
// 		IndexSingleChapter(str, &wg, mr, chapter)
// 	}

// 	wg.Wait()
// }

// func IndexSingleChapter(str *store.Store, wg *sync.WaitGroup, mr *mangareader.Mangareader, chapter string) {
// 	if mr.IsSingleIssue(chapter) {
// 		slug, chapterNumber := mr.GetInfo(chapter)

// 		result, err := str.GetMangaByMangareaderID(slug)
// 		if err != nil {
// 			mal, err := FindFromMAL(slug)
// 			if err != nil {
// 				logrus.Error(err)
// 			}

// 			if mal.MalID != 0 {
// 				anilistResult, _ := anilist.GetByMAL(strconv.FormatInt(mal.MalID, 10))
// 				var Tags []string

// 				if anilistResult == nil {
// 					logrus.Warnln("Cant Find ON Anilist", mal.MalID, mal.Title)
// 					return
// 				}

// 				for _, tag := range anilistResult.Tags {
// 					if tag.Name != nil {
// 						Tags = append(Tags, *tag.Name)
// 					}
// 				}

// 				MALID := strconv.FormatInt(mal.MalID, 10)
// 				AnilistID := strconv.FormatInt(anilistResult.ID, 10)
// 				emptyChepters := make([]primitive.ObjectID, 0)

// 				record := model.Manga{
// 					Title: mal.Title,
// 					Type:  string(mal.Type),

// 					Description:  mal.Synopsis,
// 					IsPublishing: mal.Publishing,
// 					Links: model.Links{
// 						MAL:         &MALID,
// 						Mangareader: &slug,
// 						Anilist:     &AnilistID,
// 					},
// 					Genres:   anilistResult.Genres,
// 					Tags:     Tags,
// 					Synonyms: anilistResult.Synonyms,
// 					Cover: model.Cover{
// 						Color:      anilistResult.CoverImage.Color,
// 						ExtraLarge: anilistResult.CoverImage.ExtraLarge,
// 						Large:      anilistResult.CoverImage.Large,
// 						Medium:     anilistResult.CoverImage.Medium,
// 					},
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					Banner:    anilistResult.BannerImage,
// 					StartDate: model.Date{
// 						Day:   anilistResult.StartDate.Day,
// 						Month: anilistResult.StartDate.Month,
// 						Year:  anilistResult.StartDate.Year,
// 					},
// 					EndDate: model.Date{
// 						Day:   anilistResult.EndDate.Day,
// 						Month: anilistResult.EndDate.Month,
// 						Year:  anilistResult.EndDate.Year,
// 					},
// 					Chapters: emptyChepters,
// 				}

// 				_, err := str.CreateManga(record)
// 				if err != nil {
// 					logrus.Error(err)
// 					return
// 				}

// 				result, err := str.GetMangaByMangareaderID(slug)
// 				if err != nil {
// 					logrus.Error(err)
// 					return
// 				}
// 				wg.Add(1)
// 				go IndexMangareaderChepter(str, chapterNumber, result, wg)
// 			} else {
// 				wg.Add(1)
// 				go CreateMapping(str, slug, wg)
// 			}

// 		} else {
// 			wg.Add(1)
// 			go IndexMangareaderChepter(str, chapterNumber, result, wg)
// 		}

// 	}
// }

// func IndexMangareaderChepter(str *store.Store, chepter string, record model.Manga, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// logrus.Infoln(record.Title, chepter)

// 	URL := fmt.Sprintf("https://www.mangareader.net/%s/%s/", *record.Links.Mangareader, chepter)
// 	SOURCE := "www.mangareader.net"

// 	mr := mangareader.Mangareader{}

// 	comic := new(mangareader.Comic)
// 	comic.Name = *record.Links.Mangareader
// 	comic.IssueNumber = chepter
// 	comic.URLSource = URL
// 	comic.Source = SOURCE

// 	links, err := mr.RetrieveImageLinks(comic)
// 	if err != nil {
// 		logrus.Error(err)
// 		return
// 	}
// 	if number, err := strconv.ParseInt(chepter, 10, 64); err == nil {
// 		chapterRecord := model.Chapter{
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 			Language:  "en",
// 			Number:    number,
// 			Source:    "mangareader",
// 			Links:     links,
// 			Manga:     record.ID,
// 		}

// 		str.CreateChapter(record, chapterRecord)
// 	}
// 	logrus.Infoln("INDEXED", record.Title, chepter)

// }

// func IndexMangareaderChepterWithoutWG(str *store.Store, chepter string, record model.Manga) {

// 	URL := fmt.Sprintf("https://www.mangareader.net/%s/%s/", *record.Links.Mangareader, chepter)
// 	SOURCE := "www.mangareader.net"

// 	mr := mangareader.Mangareader{}

// 	comic := new(mangareader.Comic)
// 	comic.Name = *record.Links.Mangareader
// 	comic.IssueNumber = chepter
// 	comic.URLSource = URL
// 	comic.Source = SOURCE

// 	links, err := mr.RetrieveImageLinks(comic)
// 	if err != nil {
// 		logrus.Error(err)
// 		return
// 	}
// 	if number, err := strconv.ParseInt(chepter, 10, 64); err == nil {
// 		chapterRecord := model.Chapter{
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 			Language:  "en",
// 			Number:    number,
// 			Source:    "mangareader",
// 			Links:     links,
// 			Manga:     record.ID,
// 		}

// 		str.CreateChapter(record, chapterRecord)
// 	}
// 	logrus.Infoln("INDEXED", record.Title, chepter)

// }

// func SyncAllChapters(str *store.Store) {
// 	var mr = new(mangareader.Mangareader)

// 	for _, link := range Links() {

// 		issues, err := mr.RetrieveIssueLinks("https://www.mangareader.net/"+link, false, false)
// 		if err != nil {
// 			logrus.Errorln(err)
// 		}

// 		for _, chapter := range issues {
// 			if mr.IsSingleIssue(chapter) {
// 				slug, chapterNumber := mr.GetInfo(chapter)

// 				result, err := str.GetMangaByMangareaderID(slug)
// 				if err != nil {
// 					mal, err := FindFromMAL(slug)
// 					if err != nil {
// 						logrus.Error(err)
// 					}

// 					if mal.MalID != 0 {
// 						anilistResult, _ := anilist.GetByMAL(strconv.FormatInt(mal.MalID, 10))
// 						var Tags []string

// 						if anilistResult == nil {
// 							logrus.Warnln("Cant Find ON Anilist", mal.MalID, mal.Title)
// 							return
// 						}

// 						for _, tag := range anilistResult.Tags {
// 							if tag.Name != nil {
// 								Tags = append(Tags, *tag.Name)
// 							}
// 						}

// 						MALID := strconv.FormatInt(mal.MalID, 10)
// 						AnilistID := strconv.FormatInt(anilistResult.ID, 10)
// 						emptyChepters := make([]primitive.ObjectID, 0)

// 						record := model.Manga{
// 							Title: mal.Title,
// 							Type:  string(mal.Type),

// 							Description:  mal.Synopsis,
// 							IsPublishing: mal.Publishing,
// 							Links: model.Links{
// 								MAL:         &MALID,
// 								Mangareader: &slug,
// 								Anilist:     &AnilistID,
// 							},
// 							Genres:   anilistResult.Genres,
// 							Tags:     Tags,
// 							Synonyms: anilistResult.Synonyms,
// 							Cover: model.Cover{
// 								Color:      anilistResult.CoverImage.Color,
// 								ExtraLarge: anilistResult.CoverImage.ExtraLarge,
// 								Large:      anilistResult.CoverImage.Large,
// 								Medium:     anilistResult.CoverImage.Medium,
// 							},
// 							CreatedAt: time.Now(),
// 							UpdatedAt: time.Now(),
// 							Banner:    anilistResult.BannerImage,
// 							StartDate: model.Date{
// 								Day:   anilistResult.StartDate.Day,
// 								Month: anilistResult.StartDate.Month,
// 								Year:  anilistResult.StartDate.Year,
// 							},
// 							EndDate: model.Date{
// 								Day:   anilistResult.EndDate.Day,
// 								Month: anilistResult.EndDate.Month,
// 								Year:  anilistResult.EndDate.Year,
// 							},
// 							Chapters: emptyChepters,
// 						}

// 						_, err := str.CreateManga(record)
// 						if err != nil {
// 							logrus.Error(err)
// 							return
// 						}

// 						result, err := str.GetMangaByMangareaderID(slug)
// 						if err != nil {
// 							logrus.Error(err)
// 							return
// 						}

// 						IndexMangareaderChepterWithoutWG(str, chapterNumber, result)
// 					} else {

// 						CreateMappingWithoutWG(str, slug)
// 					}

// 				} else {
// 					IndexMangareaderChepterWithoutWG(str, chapterNumber, result)
// 				}

// 			}
// 		}
// 	}

// }
