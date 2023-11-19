package controller

import (
	"bpm-wrapper/internal/data/dto"
	"bpm-wrapper/internal/usecase"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller struct {
	UseCase usecase.Usecase
	Log     *zap.SugaredLogger
	Pub     message.Publisher
}

func (c *Controller) StartProcessPassangerHandler(msg *message.Message) error {
	// var body dto.StartProcessRequest
	var body dto.StartProcessPassengerRequest

	processName := "passenger_flow"
	version := "1.0"

	err := json.Unmarshal(msg.Payload, &body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
	}

	_, err = c.UseCase.StartProcessPassenger(processName, version, body)

	if err != nil {
		c.Log.Error(err)
		msg.Ack()
	}

	return err

}

func (c *Controller) SaveWorkflow(ctx *fiber.Ctx) error {
	var body dto.SaveWorkflowRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	err = c.UseCase.SaveWorkflow(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
