package http

import (
	"github.com/gofiber/fiber/v2"
)

func GeoMappingRoutes(app *fiber.App, handler *GeoMappingHandler) {
	app.Get("/geomapping/device-list", handler.GetDevice)
	app.Post("/geomapping/device-detail", handler.GetDeviceDetail)
	app.Post("/geomapping/sensor-list", handler.GetSensor)
	app.Put("/geomapping/update-sensor", handler.UpdateSensorData)
	app.Put("/geomapping/update-device", handler.UpdateDeviceData)
	app.Get("/geomapping/group-list", handler.GetGroup)
	app.Get("/geomapping/city-list", handler.GetCity)
	app.Get("/geomapping/district-list/", handler.GetDistrict)
	app.Get("/geomapping/subdistrict-list/", handler.GetSubDistrict)
	app.Get("/geomapping/dashboard/", handler.Dashboard)
}
