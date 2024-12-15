package http

import (
	"github.com/gofiber/fiber/v2"
)

func ReportRoutes(app *fiber.App, handler *ReportHandler) {
	app.Get("/report/list", handler.ReportList)
	app.Post("/report/generate", handler.CreateReport)
	app.Post("/report/download/:id", handler.DownloadReportFile)
}
