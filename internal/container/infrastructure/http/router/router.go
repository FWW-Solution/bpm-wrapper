package router

import (
	"bpm-wrapper/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Initialize(app *fiber.App, ctrl *controller.Controller) *fiber.App {
	app.Get("/", monitor.New(monitor.Config{Title: "fww-bpm-wrapper metrics page"}))

	Api := app.Group("/api")

	_ = Api.Group("/private/v1")

	return app

}
