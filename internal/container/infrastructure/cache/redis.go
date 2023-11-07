package cache

import (
	"bpm-wrapper/internal/config"

	"github.com/redis/go-redis/v9"
)

func SetupRedis(cfg *config.RedisConfig) *redis.Client {
	connection := redis.NewClient(&redis.Options{
		Addr:            cfg.Host + ":" + cfg.Port,
		Password:        cfg.Password,
		DB:              cfg.DB,
		MaxRetries:      cfg.MaxRetries,
		PoolFIFO:        cfg.PoolFIFO,
		PoolSize:        cfg.PoolSize,
		PoolTimeout:     cfg.PoolTimeout,
		MinIdleConns:    cfg.MinIdleConns,
		MaxIdleConns:    cfg.MaxIdleConns,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
	})
	return connection
}
