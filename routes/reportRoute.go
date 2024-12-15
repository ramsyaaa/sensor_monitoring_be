package routes

import (
	"sensor_monitoring_be/modules/report/http"
	"sensor_monitoring_be/modules/report/repository"
	"sensor_monitoring_be/modules/report/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReportRouter(app *fiber.App, db *gorm.DB) {
	reportRepo := repository.NewReportRepository(db)
	reportService := service.NewReportService(reportRepo)
	reportHandler := http.NewReportHandler(reportService)

	http.ReportRoutes(app, reportHandler)
}
