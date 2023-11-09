package queue

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

func NewRouter(router *message.Router, handlerTopicExchange string, subscribeTopic string, subscriber message.Subscriber, handlerFunc func(msg *message.Message) error) (*message.Router, error) {
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

	router.AddNoPublisherHandler(
		handlerTopicExchange,
		subscribeTopic,
		subscriber,
		handlerFunc,
	)

	return router, err
}

func NewRouterWithPublisher(router *message.Router, handlerTopicExchange string, subscribeTopic string, subscriber message.Subscriber, publishTopic string, publisher message.Publisher, handlerFunc func(msg *message.Message) ([]*message.Message, error)) (*message.Router, error) {
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

	router.AddHandler(
		handlerTopicExchange,
		subscribeTopic,
		subscriber,
		publishTopic,
		publisher,
		handlerFunc,
	)

	return router, err
}
