package routes

import (
	"sensor_monitoring_be/modules/auth/http"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	authHandler := http.NewAuthHandler()
	http.AuthRoutes(app, authHandler)
}
