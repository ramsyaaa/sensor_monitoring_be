package routes

import (
	"sensor_monitoring_be/modules/cron/http"

	"github.com/gofiber/fiber/v2"
)

func CronRouter(app *fiber.App) {
	cronhHandler := http.NewCronHandler()
	http.CronRoutes(app, cronhHandler)
}
