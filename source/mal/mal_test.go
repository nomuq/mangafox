package mal

import (
	"github.com/akyoto/assert"
	"testing"
)

func TestGetManga(t *testing.T) {
	manga, err := GetManga(11)

	assert.Equal(t, int64(700), manga.Chapters)
	assert.Equal(t, "Naruto", manga.Title)
	assert.Nil(t, err)
}
