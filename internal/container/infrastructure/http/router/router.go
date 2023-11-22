package router

import (
	"bpm-wrapper/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Initialize(app *fiber.App, ctrl *controller.Controller) *fiber.App {
	app.Get("/", monitor.New(monitor.Config{Title: "fww-bpm-wrapper metrics page"}))

	Api := app.Group("/api")

	v1 := Api.Group("/private/v1")

	// bpm
	v1.Post("/workflow", ctrl.SaveWorkflow)
	v1.Post("/assign-human-task", ctrl.AssignHumanTask)

	// passanger
	v1.Put("/passenger", ctrl.UpdatePassenger)

	// payment
	v1.Post("/payment/invoice", ctrl.GenerateInvoice)
	v1.Put("/payment", ctrl.UpdatePayment)

	// ticket
	v1.Put("/ticket", ctrl.UpdateTicket)

	// notification
	v1.Post("/notification/email", ctrl.SendEmailNotification)

	return app

}
