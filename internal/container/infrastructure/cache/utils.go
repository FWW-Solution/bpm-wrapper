package cache

import (
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	// ctx         = context.Background()
)

// InitRedisClient initializes the Redis client.
func InitRedisClient(client *redis.Client) {
	redisClient = client
}
