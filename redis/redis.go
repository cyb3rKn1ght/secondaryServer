package redis

import (
	"github.com/go-redis/redis"
)

var redisInstance *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// Redis ..
func Redis() *redis.Client {
	return redisInstance
}
