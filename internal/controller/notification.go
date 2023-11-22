package controller

import (
	dtonotification "bpm-wrapper/internal/data/dto_notification"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) SendEmailNotification(ctx *fiber.Ctx) error {
	var body dtonotification.SendEmailRequest

	if err := ctx.BodyParser(&body); err != nil {
		c.Log.Error(err)
		return err
	}

	err := c.UseCase.SendEmailNotification(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON("OK")
}
