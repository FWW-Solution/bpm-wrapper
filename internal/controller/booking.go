package controller

import (
	dtobooking "bpm-wrapper/internal/data/dto_booking"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) UpdateBooking(ctx *fiber.Ctx) error {
	var body dtobooking.RequestUpdateBooking

	err := ctx.BodyParser(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	err = c.UseCase.UpdateBooking(&body)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
