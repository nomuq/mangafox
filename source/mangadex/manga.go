package mangadex

type Manga struct {
	CoverURL    string    `json:"cover_url"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	Author      string    `json:"author"`
	Status      int64     `json:"status"`
	Hentai      int64     `json:"hentai"`
	Genres      []int64   `json:"genres"`
	Language    string    `json:"lang_flag"`
	Chapters    []Chapter `json:"chapters"`
	Links       Links     `json:"links"`
}

type Links struct {
	Al  string `json:"al"`
	Ap  string `json:"ap"`
	Kt  string `json:"kt"`
	Mu  string `json:"mu"`
	Mal string `json:"mal"`
}
