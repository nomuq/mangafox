package anilist

func Search(title string) (*Manga, error) {
	type Variables struct {
		Title string `json:"title"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: `
		query ($title: String) {
			Media (search: $title, type: MANGA) {
			  id
			  title {
				romaji
				english
				native
			  } 			  
			}
		  }
		`,
		Variables: Variables{
			Title: title,
		},
	}

	// Query response
	response := new(struct {
		Data struct {
			Manga *Manga `json:"Media"`
		} `json:"data"`
	})

	err := Query(body, &response)

	if err != nil {
		return nil, err
	}

	return response.Data.Manga, nil
}
