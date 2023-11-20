package controller

import (
	dtopayment "bpm-wrapper/internal/data/dto_payment"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GenerateInvoice(ctx *fiber.Ctx) error {
	var body dtopayment.GenerateInvoiceRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	err = c.UseCase.GenerateInvoice(body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (c *Controller) DoPaymentHandler(msg *message.Message) error {
	var body dtopayment.DoPaymentRequest

	err := json.Unmarshal(msg.Payload, &body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
	}

	err = c.UseCase.DoPayment(&body)

	if err != nil {
		c.Log.Error(err)
		msg.Ack()
	}

	return err
}

func (c *Controller) UpdatePayment(ctx *fiber.Ctx) error {
	var body dtopayment.RequestUpdatePayment

	err := ctx.BodyParser(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	err = c.UseCase.UpdatePayment(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
