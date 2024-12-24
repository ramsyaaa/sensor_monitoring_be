package routes

import (
	"sensor_monitoring_be/modules/auth/http"
	"sensor_monitoring_be/modules/auth/repository"
	"sensor_monitoring_be/modules/auth/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(app *fiber.App, db *gorm.DB) {
	geomappingRepo := repository.NewAuthRepository(db)
	geomappingService := service.NewAuthService(geomappingRepo)
	geomappingHandler := http.NewAuthHandler(geomappingService)

	http.AuthRoutes(app, geomappingHandler)
}
