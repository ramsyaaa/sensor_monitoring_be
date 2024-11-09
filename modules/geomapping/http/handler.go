package http

import (
	"context"
	"net/http"

	"sensor_monitoring_be/helper"
	"sensor_monitoring_be/modules/geomapping/service"

	"github.com/gofiber/fiber/v2"
)

type GeoMappingHandler struct {
	service service.GeoMappingService
}

func NewGeoMappingHandler(service service.GeoMappingService) *GeoMappingHandler {
	return &GeoMappingHandler{service: service}
}

func (h *GeoMappingHandler) GetDevice(c *fiber.Ctx) error {
	ctx := context.Background()
	devices, err := h.service.GetDevice(ctx)
	if err != nil {
		response := helper.APIResponse("Failed to fetch device data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("Device data fetched successfully", http.StatusOK, "OK", devices)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GeoMappingHandler) GetSensor(c *fiber.Ctx) error {
	ctx := context.Background()
	type Request struct {
		DeviceID int `json:"deviceId"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid device id", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	sensor, err := h.service.GetSensor(ctx, req.DeviceID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch sensor data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Sensor data fetched successfully", http.StatusOK, "OK", sensor)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GeoMappingHandler) UpdateSensorData(c *fiber.Ctx) error {
	ctx := context.Background()
	type Request struct {
		SensorID int    `json:"sensorId"`
		Lat      string `json:"lat"`
		Lng      string `json:"lng"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid sensor data", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	data := map[string]interface{}{
		"lat": req.Lat,
		"lng": req.Lng,
	}

	err := h.service.UpdateSensorData(ctx, req.SensorID, data)
	if err != nil {
		response := helper.APIResponse("Failed to update sensor data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Sensor data updated successfully", http.StatusOK, "OK", nil)
	return c.Status(http.StatusOK).JSON(response)
}
