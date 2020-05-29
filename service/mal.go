package service

import (
	"strings"

	"github.com/gosimple/slug"
	"github.com/nokusukun/jikan2go/manga"
)

func FindFromMAL(title string) (manga.Result, error) {
	result, err := manga.Search(manga.Query{Q: title})

	if err != nil {
		return manga.Result{}, err
	}

	for _, r := range result.Results {
		if strings.ToUpper(slug.Make(r.Title)) == strings.ToUpper(title) {
			return r, err
		}
	}
	return manga.Result{}, err
}
