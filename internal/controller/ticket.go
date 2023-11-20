package controller

import (
	dtoticket "bpm-wrapper/internal/data/dto_ticket"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) UpdateTicket(ctx *fiber.Ctx) error {
	var body dtoticket.RequestUpdateTicket

	err := ctx.BodyParser(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	err = c.UseCase.UpdateTicket(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (c *Controller) RedeemTicketHandler(msg *message.Message) error {
	var body dtoticket.RequestRedeemTicketToBPM

	err := json.Unmarshal(msg.Payload, &body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	err = c.UseCase.RedeemTicket(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return nil
}
