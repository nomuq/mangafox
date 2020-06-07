package anilist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	manga, err := GetByID("30011")

	assert.Equal(t, int64(700), *manga.Chapters)
	assert.Equal(t, "Naruto", manga.Title.English)
	assert.Nil(t, err)
}
