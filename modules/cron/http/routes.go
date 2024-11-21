package http

import (
	"github.com/gofiber/fiber/v2"
)

func CronRoutes(app *fiber.App, handler *CronHandler) {
	app.Post("/fetch-data", handler.FetchData)
}
