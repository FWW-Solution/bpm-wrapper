package controller

import (
	dtopassenger "bpm-wrapper/internal/data/dto_passenger"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) UpdatePassenger(ctx *fiber.Ctx) error {
	var body dtopassenger.Passenger

	err := ctx.BodyParser(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	err = c.UseCase.UpdatePassenger(body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
