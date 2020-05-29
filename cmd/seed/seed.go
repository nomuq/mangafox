package main

import (
	"context"
	"mangafox/store"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	// md := mangadex.New()
// 	// m, _, err := md.Manga("9967")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// for _, chapter := range m.Chapters {
// 	// 	fmt.Println(chapter.Pages)
// 	// }
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 24*time.Hour)
// 	err = client.Connect(ctx)

// 	if err != nil {
// 		panic(err)
// 	}

// 	collection := client.Database("mangadex").Collection("manga")

// 	// mangas, err := str.GetAllManga()
// 	// if err != nil {
// 	// 	logrus.Fatalln(err)
// 	// }

// 	// data, _ := json.Marshal(mangas)
// 	// ioutil.WriteFile("mangas.json", data, os.ModePerm)
// 	collection := client.Database("mal").Collection("manga")
// 	for i := 1; i < 99999; i++ {
// 		url := "https://mangadex.org/api/manga/" + i
// 		method := "GET"

// 		client := &http.Client{}
// 		req, err := http.NewRequest(method, url, nil)

// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		res, err := client.Do(req)
// 		defer res.Body.Close()
// 		body, err := ioutil.ReadAll(res.Body)
// 		if err == nil {
// 			res, err := collection.InsertOne(ctx, bolB)
// 			if err == nil {
// 				fmt.Println(res.InsertedID)
// 			} else {
// 				fmt.Println(err)
// 			}
// 		}

// 	}

// 	str.Client.Disconnect(ctx)

// }

// type Result struct {
// 	Manga   Manga              `json:"manga"`
// 	Chapter map[string]Chapter `json:"chapter"`
// 	Group   map[string]Group   `json:"group"`
// 	Status  string             `json:"status"`
// }

// type Chapter struct {
// 	Volume     string      `json:"volume"`
// 	Chapter    string      `json:"chapter"`
// 	Title      string      `json:"title"`
// 	LangCode   LangCode    `json:"lang_code"`
// 	GroupID    int64       `json:"group_id"`
// 	GroupName  GroupName   `json:"group_name"`
// 	GroupID2   int64       `json:"group_id_2"`
// 	GroupName2 interface{} `json:"group_name_2"`
// 	GroupID3   int64       `json:"group_id_3"`
// 	GroupName3 interface{} `json:"group_name_3"`
// 	Timestamp  int64       `json:"timestamp"`
// }

// type Group struct {
// 	GroupName GroupName `json:"group_name"`
// }

// type Manga struct {
// 	CoverURL    string  `json:"cover_url"`
// 	Description string  `json:"description"`
// 	Title       string  `json:"title"`
// 	Artist      string  `json:"artist"`
// 	Author      string  `json:"author"`
// 	Status      int64   `json:"status"`
// 	Genres      []int64 `json:"genres"`
// 	LastChapter string  `json:"last_chapter"`
// 	LangName    string  `json:"lang_name"`
// 	LangFlag    string  `json:"lang_flag"`
// 	Hentai      int64   `json:"hentai"`
// 	Links       Links   `json:"links"`
// }

// type Links struct {
// 	Al    string `json:"al"`
// 	Ap    string `json:"ap"`
// 	BW    string `json:"bw"`
// 	Kt    string `json:"kt"`
// 	Mu    string `json:"mu"`
// 	Amz   string `json:"amz"`
// 	Ebj   string `json:"ebj"`
// 	Mal   string `json:"mal"`
// 	Raw   string `json:"raw"`
// 	Engtl string `json:"engtl"`
// }

// type GroupName string

// const (
// 	AnimeProDestiny       GroupName = "Anime Pro Destiny"
// 	CandyCaneTranslations GroupName = "Candy Cane Translations"
// 	ChudoManga            GroupName = "Chudo Manga"
// 	CrunchyrollExLicenses GroupName = "Crunchyroll (Ex-Licenses)"
// 	FBIKun                GroupName = "FBI-kun"
// 	GekkouScans           GroupName = "Gekkou Scans"
// 	TheCoordinate         GroupName = "The Coordinate"
// 	WhiteoutScans         GroupName = "Whiteout Scans"
// )

// type LangCode string

// const (
// 	Br LangCode = "br"
// 	GB LangCode = "gb"
// 	It LangCode = "it"
// 	Ph LangCode = "ph"
// 	Ru LangCode = "ru"
// )

func main() {
	ctx := context.Background()
	str, err := store.New(ctx, "mongodb://localhost:27017")
	if err != nil {
		logrus.Fatalln(err)
	}
	CreateIndexes(ctx, str)
	ctx.Done()
}

func CreateIndexes(ctx context.Context, store *store.Store) {

	collection := store.MangaCollection()

	indexes := []mongo.IndexModel{
		mongo.IndexModel{
			Keys: bson.M{
				"links.anilist": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mal": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mangadex": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mangareader": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"links.mangatown": 1,
			},
		},
		mongo.IndexModel{
			Keys: bson.M{
				"isPublishing": 1,
			},
		},
	}

	res, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Infoln(res)
}
