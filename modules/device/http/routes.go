package http

import (
	"github.com/gofiber/fiber/v2"
)

func DeviceRoutes(app *fiber.App, handler *DeviceHandler) {
	deviceHandler := NewDeviceHandler()

	app.Post("/get-device", deviceHandler.HandleGetDevices)
	app.Post("/get-single-device", deviceHandler.HandleGetSingleDevice)
	app.Post("/get-single-sensor", deviceHandler.HandleGetSingleSensor)
	app.Post("/get-sensor-history", deviceHandler.HandleGetSensorHistory)
	app.Post("/get-device/with-sensor", deviceHandler.HandleGetDevicesWithSensor)
}
