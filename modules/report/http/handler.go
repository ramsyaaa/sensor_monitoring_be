package http

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"sensor_monitoring_be/helper"
	"sensor_monitoring_be/models"
	"sensor_monitoring_be/modules/report/service"

	"github.com/gofiber/fiber/v2"

	"github.com/xuri/excelize/v2"
)

type ReportHandler struct {
	reportService service.ReportService
}

func NewReportHandler(reportService service.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

func (h *ReportHandler) CreateReport(c *fiber.Ctx) error {
	var data models.GeneratedReport
	if err := c.BodyParser(&data); err != nil {
		response := helper.APIResponse("Invalid request body", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	ctx := context.Background()
	report, err := h.reportService.CreateReport(ctx, data)
	if err != nil {
		response := helper.APIResponse("Failed to create report", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Return success response
	response := helper.APIResponse("Report creation initiated successfully", http.StatusOK, "OK", report)
	c.Status(http.StatusOK).JSON(response)

	go h.ExportSensorDataToExcel(int64(data.DeviceID), data.StartDate, data.EndDate, ctx)
	return nil
}

func (h *ReportHandler) ExportSensorDataToExcel(deviceID int64, startDate string, endDate string, ctx context.Context) {
	sensors, err := h.reportService.GetSensor(ctx, deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}

	var sensorData []map[string]interface{}
	for _, sensor := range sensors {
		sensorID, ok := sensor["id"].(int64)
		if !ok {
			continue
		}
		data, err := h.reportService.GetSensorData(ctx, sensorID, startDate, endDate)

		if err != nil {
			continue
		}
		sensorData = append(sensorData, data...)
	}

	f := excelize.NewFile()

	// Ambil deviceID dari data pertama
	deviceID = sensorData[0]["device_id"].(int64)

	// Mengelompokkan data berdasarkan sensor_name
	sensorGroups := make(map[string][]map[string]interface{})
	for _, data := range sensorData {
		sensorName := data["sensor_name"].(string)
		sensorGroups[sensorName] = append(sensorGroups[sensorName], data)
	}

	// Untuk setiap sensor, buat sheet dan isi datanya
	for sensorName, dataGroup := range sensorGroups {
		sheetName := sensorName
		f.NewSheet(sheetName)

		// Header
		f.SetCellValue(sheetName, "A1", "Date")
		f.SetCellValue(sheetName, "B1", "Value")

		// Set Date cell width to 145 pixels
		f.SetColWidth(sheetName, "A", "A", 20)

		// Isi data
		for i, data := range dataGroup {
			row := i + 2
			date := data["date"]
			value := data["value"].(float64)

			f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), date)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), value)
		}
	}
	f.DeleteSheet("Sheet1") // Hapus sheet default yang bernama "Sheet1"

	// Simpan file
	t := time.Now()
	filename := fmt.Sprintf("report/generated/Export-%d-%02d-%02dT%02d-%02d-%02d.xlsx",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
		return
	}

	// Update report status
	if err := h.reportService.UpdateReport(ctx, deviceID, startDate, endDate, "generated", filename); err != nil {
		fmt.Println(err)
		return
	}
}

func (h *ReportHandler) ReportList(c *fiber.Ctx) error {
	ctx := context.Background()
	reports, err := h.reportService.ReportList(ctx)
	if err != nil {
		response := helper.APIResponse("Failed to get report list data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Report list data retrieved successfully", http.StatusOK, "OK", reports)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *ReportHandler) DownloadReportFile(c *fiber.Ctx) error {
	ctx := context.Background()
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		response := helper.APIResponse("Failed to download report file", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	filename, err := h.reportService.DownloadReportFile(ctx, id)
	if err != nil {
		response := helper.APIResponse("Failed to download report file", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	if filename == "" {
		response := helper.APIResponse("Report file not found", http.StatusNotFound, "ERROR", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	filenameWithoutPath := filepath.Base(filename)
	return c.Download(filename, strings.TrimPrefix(filenameWithoutPath, "report/generated/"))
}
