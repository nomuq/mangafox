package anilist

import (
	"fmt"
	"github.com/aerogo/http/client"
	"net/http"
)

var headers = client.Headers{
	"User-Agent":   "manga-community",
	"Content-Type": "application/json",
}

// Query queries the AniList GraphQL API.
func Query(body interface{}, target interface{}) error {
	response, err := client.Post("https://graphql.anilist.co").Headers(headers).BodyJSON(body).EndStruct(target)

	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("Status: %d\n%s", response.StatusCode(), response.String())
	}

	return nil
}
