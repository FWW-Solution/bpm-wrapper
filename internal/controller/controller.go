package controller

import (
	"bpm-wrapper/internal/data/dto"
	"bpm-wrapper/internal/usecase"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
)

type Controller struct {
	UseCase usecase.Usecase
	log     *zap.SugaredLogger
	pub     message.Publisher
}

func (c *Controller) StartProcess(msg *message.Message) error {
	// var body dto.StartProcessRequest
	var body dto.StartProcessRequest

	err := json.Unmarshal(msg.Payload, &body)
	if err != nil {
		log.Fatal(err)
	}

	result, err := c.UseCase.StartProcess(body.Version)

	if err != nil {
		log.Println(err)
		msg.Ack()
	} else {
		c.pub.Publish("process_started", message.NewMessage(msg.UUID, []byte(result)))
		msg.Ack()
	}

	return err

}
