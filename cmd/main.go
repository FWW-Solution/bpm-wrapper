package main

import (
	"bpm-wrapper/internal/container"
	"bpm-wrapper/internal/container/infrastructure/http"
	"context"
	"log"

	"bpm-wrapper/internal/config"

	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	cfg := config.InitConfig()
	app, messageRouters := container.InitService(cfg)

	for _, router := range messageRouters {
		ctx := context.Background()
		go func(router message.Router) {
			err := router.Run(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}(router)
	}

	// run service
	http.StartHttpServer(app, cfg.HttpServer.Port)
}
