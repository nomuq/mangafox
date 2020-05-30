package cache

import (
	"github.com/akyoto/assert"
	"strconv"
	"testing"
)

func TestCache_GetBool(t *testing.T) {
	cache := Cache{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
	}
	err := cache.Connect()

	TestKey2, err := cache.Set("TestKey", true)
	TestKey2, err = cache.Get("TestKey")

	boolean, err := strconv.ParseBool(TestKey2)

	assert.Equal(t, boolean, true)
	assert.Nil(t, err)
}
