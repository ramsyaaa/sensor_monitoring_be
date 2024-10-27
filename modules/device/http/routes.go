package http

import (
	"github.com/gofiber/fiber/v2"
)

func DeviceRoutes(app *fiber.App, handler *DeviceHandler) {
	deviceHandler := NewDeviceHandler()

	app.Get("/get-device", deviceHandler.HandleGetDevices)
}
