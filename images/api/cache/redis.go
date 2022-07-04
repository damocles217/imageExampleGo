package cache

import (
	"github.com/go-redis/redis/v8"
)

func CreateCacheServer() *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.65:6379",
		DB:       0,
		Password: "",
	})

	return cache
}
