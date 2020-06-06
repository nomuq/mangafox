package mangadex

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestMangadex_GetInfo(t *testing.T) {
	md := Initilize()
	manga, err := md.GetInfo("5")

	assert.Equal(t, int64(2), manga.Status)
	assert.Equal(t, "Naruto", manga.Title)
	assert.Nil(t, err)
}
