package routes

import (
	"sensor_monitoring_be/modules/device/http"

	"github.com/gofiber/fiber/v2"
)

func DeviceRouter(app *fiber.App) {
	deviceHandler := http.NewDeviceHandler()
	http.DeviceRoutes(app, deviceHandler)
}
