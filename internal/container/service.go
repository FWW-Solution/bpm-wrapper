package container

import (
	"bpm-wrapper/internal/adapter"
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/container/infrastructure/cache"
	"bpm-wrapper/internal/container/infrastructure/database"
	"bpm-wrapper/internal/container/infrastructure/http"
	"bpm-wrapper/internal/container/infrastructure/http/router"
	httpclient "bpm-wrapper/internal/container/infrastructure/http_client"
	logger "bpm-wrapper/internal/container/infrastructure/log"
	"bpm-wrapper/internal/container/infrastructure/queue"
	"bpm-wrapper/internal/controller"
	"bpm-wrapper/internal/repository"
	"bpm-wrapper/internal/usecase"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
)

func InitService(cfg *config.Config) (*fiber.App, []*message.Router) {

	// Init DB
	db := database.GetConnection(&cfg.Database)
	// init redis
	clientRedis := cache.SetupRedis(&cfg.Cache)
	// init redis cache
	cache.InitRedisClient(clientRedis)
	// Init Logger
	log := logger.Initialize(cfg)
	// Init HTTP Server
	server := http.SetupHttpEngine()

	amqpMessageStream := queue.NewAmpq(&cfg.Queue)
	// set message stream subscriber
	sub, err := amqpMessageStream.NewSubscriber()
	if err != nil {
		panic(err)
	}

	// set message stream publisher
	pub, err := amqpMessageStream.NewPublisher()
	if err != nil {
		panic(err)
	}

	// Init Circuit Breaker
	cb := httpclient.InitCircuitBreaker(&cfg.HttpClient, "consecutive")
	// Init HTTP Client
	client := httpclient.InitHttpClient(&cfg.HttpClient, cb)

	// Init Adapter
	adapter := adapter.New(client, &cfg.Bonita)

	// Init Repository
	repo := repository.New(db)

	// Init UseCase
	usecase := usecase.New(&adapter, &cfg.Bonita, clientRedis, pub, repo)

	// Init Controller
	ctrl := controller.Controller{UseCase: usecase, Log: log, Pub: pub}

	// Init router
	var messageRouters []*message.Router

	startProcessPassangerRouter, err := queue.NewRouter(pub, "start_process_passanger_poison", "start_process_passanger", "start_process_passanger", sub, ctrl.StartProcessPassangerHandler)
	if err != nil {
		log.Fatal(err)
	}

	startProcessBookingRouter, err := queue.NewRouter(pub, "start_process_booking_poison", "start_process_booking", "start_process_booking", sub, ctrl.StartProcessBookingHandler)
	if err != nil {
		log.Fatal(err)
	}

	doPaymentRouter, err := queue.NewRouter(pub, "do_payment_bpm_poison", "do_payment_bpm_handler", "do_payment_bpm", sub, ctrl.DoPaymentHandler)
	if err != nil {
		log.Fatal(err)
	}

	messageRouters = append(messageRouters, startProcessPassangerRouter, startProcessBookingRouter, doPaymentRouter)

	// Init Router
	app := router.Initialize(server, &ctrl)

	return app, messageRouters
}
