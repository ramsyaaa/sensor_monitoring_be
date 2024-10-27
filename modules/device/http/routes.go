package http

import (
	"github.com/gofiber/fiber/v2"
)

func DeviceRoutes(app *fiber.App, handler *DeviceHandler) {
	deviceHandler := NewDeviceHandler()

	app.Post("/get-device", deviceHandler.HandleGetDevices)
	app.Post("/get-device/with-sensor", deviceHandler.HandleGetDevicesWithSensor)
}
