package cache

import (
	"strconv"

	"github.com/go-redis/redis/v7"
)

type Cache struct {
	Address  string
	Password string
	DB       int
	redis    *redis.Client
}

func (cache *Cache) Connect() error {
	redis := redis.NewClient(&redis.Options{
		Addr:     cache.Address,
		Password: cache.Password,
		DB:       cache.DB,
	})
	_, err := redis.Ping().Result()
	if err != nil {
		return err
	}
	cache.redis = redis
	return nil
}

func (cache Cache) Get(key string) (string, error) {
	result, err := cache.redis.Get(key).Result()
	return result, err
}

func (cache Cache) Set(key string, value interface{}) (string, error) {
	result, err := cache.redis.Set(key, value, 0).Result()
	return result, err
}

func (cache Cache) GetBool(key string) (bool, error) {
	return false, nil
	result, err := cache.redis.Get(key).Result()
	if err != nil {
		return false, err
	}

	boolean, err := strconv.ParseBool(result)
	if err != nil {
		return false, err
	}
	return boolean, err
}

func (cache Cache) SetBool(key string, value bool) (bool, error) {
	result, err := cache.redis.Set(key, strconv.FormatBool(value), 0).Result()
	if err != nil {
		return false, err
	}

	boolean, err := strconv.ParseBool(result)
	if err != nil {
		return false, err
	}
	return boolean, err
}
