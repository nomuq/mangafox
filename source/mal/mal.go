package mal

import (
	"github.com/nokusukun/jikan2go/manga"
)

func GetManga(id int64) (manga.Manga, error) {
	return manga.GetManga(manga.Manga{MalID: id})
}
