package container

import (
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/container/infrastructure/cache"
	"bpm-wrapper/internal/container/infrastructure/queue"
)

func InitService(cfg *config.Config) {

	// init redis
	clientRedis := cache.SetupRedis(cfg.Cache)
	// init redis cache
	cache.InitRedisClient(clientRedis)
	// set message stream subscriber
	_, err := queue.NewSubscriber(cfg.Queue)
	if err != nil {
		panic(err)
	}

	// set message stream publisher
	_, err = queue.NewPublisher(cfg.Queue)
	if err != nil {
		panic(err)
	}
}
