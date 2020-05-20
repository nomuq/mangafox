package anilist

type Manga struct {
	ID              int64       `json:"id,omitempty"`
	IDMal           *int64      `json:"idMal,omitempty"`
	Title           *Title      `json:"title,omitempty"`
	Genres          []string    `json:"genres"`
	Synonyms        []string    `json:"synonyms"`
	Description     *string     `json:"description,omitempty"`
	Type            *string     `json:"type,omitempty"`
	Format          *string     `json:"format,omitempty"`
	StartDate       *Date       `json:"startDate,omitempty"`
	EndDate         *Date       `json:"endDate,omitempty"`
	UpdatedAt       *int64      `json:"updatedAt,omitempty"`
	CoverImage      *CoverImage `json:"coverImage,omitempty"`
	BannerImage     *string     `json:"bannerImage,omitempty"`
	Tags            []Tag       `json:"tags"`
	Status          *string     `json:"status,omitempty"`
	Chapters        *int64      `json:"chapters"`
	Popularity      *int64      `json:"popularity,omitempty"`
	CountryOfOrigin *string     `json:"countryOfOrigin,omitempty"`
	Staff           *Staff      `json:"staff,omitempty"`
	Characters      *Characters `json:"characters,omitempty"`
}

type Characters struct {
	Edges []CharactersEdge `json:"edges"`
}

type CharactersEdge struct {
	ID   *int64      `json:"id,omitempty"`
	Role *string     `json:"role,omitempty"`
	Node *PurpleNode `json:"node,omitempty"`
}

type PurpleNode struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *Name   `json:"name,omitempty"`
	Image       *Image  `json:"image,omitempty"`
	Description *string `json:"description"`
}

type Image struct {
	Large  *string `json:"large,omitempty"`
	Medium *string `json:"medium,omitempty"`
}

type Name struct {
	First  *string `json:"first,omitempty"`
	Last   *string `json:"last"`
	Full   *string `json:"full,omitempty"`
	Native *string `json:"native,omitempty"`
}

type CoverImage struct {
	ExtraLarge *string `json:"extraLarge,omitempty"`
	Large      *string `json:"large,omitempty"`
	Medium     *string `json:"medium,omitempty"`
	Color      *string `json:"color"`
}

type Date struct {
	Year  *int `json:"year"`
	Month *int `json:"month"`
	Day   *int `json:"day"`
}

type Staff struct {
	Edges []StaffEdge `json:"edges"`
}

type StaffEdge struct {
	ID   *int64      `json:"id,omitempty"`
	Role *string     `json:"role,omitempty"`
	Node *FluffyNode `json:"node,omitempty"`
}

type FluffyNode struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *Name   `json:"name,omitempty"`
	Language    *string `json:"language,omitempty"`
	Image       *Image  `json:"image,omitempty"`
	Description *string `json:"description"`
}

type Tag struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Category    *string `json:"category,omitempty"`
	Rank        *int64  `json:"rank,omitempty"`
}

type Title struct {
	Romaji  *string `json:"romaji,omitempty"`
	English *string `json:"english"`
	Native  *string `json:"native,omitempty"`
}

// Link returns the permalink to that manga.
// func (manga *Manga) Link() string {
// 	return "https://anilist.co/manga/" + strconv.Itoa(manga.ID)
// }

func GetByMAL(idMal string) (*Manga, error) {
	type Variables struct {
		IDMal string `json:"idMal"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: `
		

query ($idMal: Int) { # Define which variables will be used in the query (id)
	Media (idMal: $idMal, type: MANGA) { # Insert our variables into the query arguments (id) (type: ANIME is hard-coded in the query)
	  id
	  title {
		romaji
		english
		native
	  }     
	  genres
	  synonyms
	  description
	  type
	  format
	  startDate {
		year
		month
		day
	  }
	  endDate {
		year
		month
		day
	  }
	  updatedAt
	  coverImage {
		extraLarge
		large
		medium
		color
	  }
	  bannerImage
	  tags {
		id
		name
		description
		category
		rank      
	  }
	  status
	  chapters
	  popularity
	  countryOfOrigin
	  staff {
		edges {
		  id
		  role
		  node {
			id
			name {
			  first
			  last
			  full
			  native
			}
			language
			image {
			  large
			  medium
			}
			description
			
		  }
		}
	  }
	  characters {
		edges {
		  id
		  role
		  node {
			id
			name {
			  first
			  last
			  full
			  native
			}
			image {
			  large
			  medium
			}
			description
			
		  }
		  
		}
	  }
	}
  }
  
  
		`,
		Variables: Variables{
			IDMal: idMal,
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
