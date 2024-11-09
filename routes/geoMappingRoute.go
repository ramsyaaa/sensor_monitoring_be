package routes

import (
	"sensor_monitoring_be/modules/geomapping/http"
	"sensor_monitoring_be/modules/geomapping/repository"
	"sensor_monitoring_be/modules/geomapping/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GeoMappingRouter(app *fiber.App, db *gorm.DB) {
	geomappingRepo := repository.NewGeoMappingRepository(db)
	geomappingService := service.NewGeoMappingService(geomappingRepo)
	geomappingHandler := http.NewGeoMappingHandler(geomappingService)

	http.GeoMappingRoutes(app, geomappingHandler)
}
