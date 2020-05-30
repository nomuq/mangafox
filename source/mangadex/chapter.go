package mangadex

type Chapter struct {
	Title    string   `json:"title"`
	Language string   `json:"lang_code"`
	Number   string   `json:"chapter"`
	ID       int64    `json:"id"`
	Hash     string   `json:"hash"`
	MangaID  int64    `json:"manga_id"`
	Server   string   `json:"server"`
	Links    []string `json:"page_array"`
}
