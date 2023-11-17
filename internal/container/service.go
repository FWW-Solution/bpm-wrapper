package container

import (
	"bpm-wrapper/internal/adapter"
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/container/infrastructure/cache"
	"bpm-wrapper/internal/container/infrastructure/http"
	"bpm-wrapper/internal/container/infrastructure/http/router"
	logger "bpm-wrapper/internal/container/infrastructure/log"
	"bpm-wrapper/internal/container/infrastructure/queue"
	"bpm-wrapper/internal/controller"
	"bpm-wrapper/internal/usecase"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
)

func InitService(cfg *config.Config) (*fiber.App, []message.Router) {

	// init redis
	clientRedis := cache.SetupRedis(cfg.Cache)
	// init redis cache
	cache.InitRedisClient(clientRedis)
	// Init Logger
	log := logger.Initialize(cfg)
	// Init HTTP Server
	server := http.SetupHttpEngine()
	// set message stream subscriber
	sub, err := queue.NewSubscriber(cfg.Queue)
	if err != nil {
		panic(err)
	}

	// set message stream publisher
	pub, err := queue.NewPublisher(cfg.Queue)
	if err != nil {
		panic(err)
	}

	// Init Adapter
	adapter := adapter.New(nil, cfg.Bonita)

	// Init UseCase
	usecase := usecase.New(&adapter, cfg.Bonita, clientRedis)

	// Init Controller
	ctrl := controller.Controller{UseCase: usecase, Log: log, Pub: pub}

	// Init router
	var messageRouters []message.Router

	startProcessPassangerRouter, err := queue.NewRouter(pub, "start_process_passanger_poison", "start_process_passanger", "start_process_passanger", sub, ctrl.StartProcessPassangerHandler)
	if err != nil {
		log.Fatal(err)
	}

	messageRouters = append(messageRouters, *startProcessPassangerRouter)

	// Init Router
	app := router.Initialize(server, &ctrl)

	return app, messageRouters
}
