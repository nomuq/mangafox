package mal

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestGetManga(t *testing.T) {
	manga, err := GetManga(11)

	assert.Equal(t, int64(700), manga.Chapters)
	assert.Equal(t, "Naruto", manga.Title)
	assert.Nil(t, err)
}
