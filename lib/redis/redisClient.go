package redis

import (
	"github.com/go-redis/redis/v8"
	"os"
)

func Client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})
	return rdb
}
