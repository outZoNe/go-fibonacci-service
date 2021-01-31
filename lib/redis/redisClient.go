package redis

import (
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
)

var redisClient *redis.Client
var once sync.Once

func Client() *redis.Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR") + ":" + os.Getenv("REDIS_PORT"),
			Password: "",
			DB:       0,
		})
	})

	return redisClient
}
