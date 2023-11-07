package queue

import (
	"bpm-wrapper/internal/config"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

func NewSubscriber(cfg *config.MessageQueueConfig) (message.Subscriber, error) {
	ampqURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	ampqConfig := amqp.NewDurableQueueConfig(ampqURI)

	subscriber, err := amqp.NewSubscriber(
		ampqConfig,
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		log.Fatal(err)
	}

	return subscriber, err
}

func NewPublisher(cfg *config.MessageQueueConfig) (message.Publisher, error) {
	ampqURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	ampqConfig := amqp.NewDurableQueueConfig(ampqURI)

	publisher, err := amqp.NewPublisher(
		ampqConfig,
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		log.Fatal(err)
	}

	return publisher, err
}

func ProcessMessages(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("Got message: %s", string(msg.Payload))
		msg.Ack()
	}
}

func NewRouter() (*message.Router, error) {
	logger := watermill.NewStdLogger(true, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		log.Fatal(err)
	}

	router.AddPlugin(plugin.SignalsHandler)

	router.AddMiddleware(
		middleware.CorrelationID,

		middleware.Retry{
			MaxRetries:      5,
			InitialInterval: 500,
			Logger:          logger,
		}.Middleware,

		middleware.Recoverer,
	)

	return router, err
}
