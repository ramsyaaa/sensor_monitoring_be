package http

import (
	"github.com/gofiber/fiber/v2"
)

func GeoMappingRoutes(app *fiber.App, handler *GeoMappingHandler) {
	app.Get("/geomapping/device-list", handler.GetDevice)
	app.Post("/geomapping/sensor-list", handler.GetSensor)
	app.Put("/geomapping/update-sensor", handler.UpdateSensorData)
}
