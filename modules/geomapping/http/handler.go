package http

import (
	"context"
	"net/http"
	"strconv"

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
	groupID, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		groupID = 0
	}
	cityID, err := strconv.Atoi(c.Query("city_id"))
	if err != nil {
		cityID = 0
	}
	districtID, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		districtID = 0
	}
	subdistrictID, err := strconv.Atoi(c.Query("subdistrict_id"))
	if err != nil {
		subdistrictID = 0
	}
	keyword := c.Query("keyword")

	devices, err := h.service.GetDevice(ctx, groupID, cityID, districtID, subdistrictID, keyword)
	if err != nil {
		response := helper.APIResponse("Failed to fetch device data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("Device data fetched successfully", http.StatusOK, "OK", devices)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GeoMappingHandler) GetDeviceDetail(c *fiber.Ctx) error {
	ctx := context.Background()
	type Request struct {
		DeviceID int `json:"deviceId"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid device id", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	device, err := h.service.GetDeviceDetail(ctx, req.DeviceID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch device data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Device Detail data fetched successfully", http.StatusOK, "OK", device)
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

func (h *GeoMappingHandler) UpdateDeviceData(c *fiber.Ctx) error {
	ctx := context.Background()
	type Request struct {
		DeviceID            int    `json:"device_id"`
		Lat                 string `json:"lat"`
		Lng                 string `json:"lng"`
		CityID              int    `json:"city_id"`
		DistrictID          int    `json:"district_id"`
		SubdistrictID       int    `json:"subdistrict_id"`
		PointCode           string `json:"point_code"`
		Address             string `json:"address"`
		ElectricalPanel     string `json:"electrical_panel"`
		SurroundingWaters   string `json:"surrounding_waters"`
		LocationInformation string `json:"location_information"`
		Note                string `json:"note"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid request data", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if req.DeviceID == 0 {
		response := helper.APIResponse("Invalid device id", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	data := map[string]interface{}{
		"city_id":              req.CityID,
		"lat":                  req.Lat,
		"lng":                  req.Lng,
		"district_id":          req.DistrictID,
		"subdistrict_id":       req.SubdistrictID,
		"point_code":           req.PointCode,
		"address":              req.Address,
		"electrical_panel":     req.ElectricalPanel,
		"surrounding_waters":   req.SurroundingWaters,
		"location_information": req.LocationInformation,
		"note":                 req.Note,
	}

	err := h.service.UpdateDeviceData(ctx, req.DeviceID, data)
	if err != nil {
		response := helper.APIResponse("Failed to update device data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Device data updated successfully", http.StatusOK, "OK", nil)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GeoMappingHandler) GetGroup(c *fiber.Ctx) error {
	ctx := context.Background()
	groups, err := h.service.GetGroup(ctx)
	if err != nil {
		response := helper.APIResponse("Failed to get group data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Group data retrieved successfully", http.StatusOK, "OK", groups)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GeoMappingHandler) GetCity(c *fiber.Ctx) error {
	ctx := context.Background()
	cities, err := h.service.GetCity(ctx)
	if err != nil {
		response := helper.APIResponse("Failed to get city data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("City data retrieved successfully", http.StatusOK, "OK", cities)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GeoMappingHandler) GetDistrict(c *fiber.Ctx) error {
	ctx := context.Background()
	cityId, err := strconv.Atoi(c.Query("city_id"))
	if err != nil {
		response := helper.APIResponse("Invalid city id", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	districts, err := h.service.GetDistrict(ctx, cityId)
	if err != nil {
		response := helper.APIResponse("Failed to get district data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("District data retrieved successfully", http.StatusOK, "OK", districts)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GeoMappingHandler) GetSubDistrict(c *fiber.Ctx) error {
	ctx := context.Background()
	districtId, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		response := helper.APIResponse("Invalid district id", http.StatusBadRequest, "ERROR", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	subDistricts, err := h.service.GetSubDistrict(ctx, districtId)
	if err != nil {
		response := helper.APIResponse("Failed to get subdistrict data", http.StatusInternalServerError, "ERROR", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("Subdistrict data retrieved successfully", http.StatusOK, "OK", subDistricts)
	return c.Status(http.StatusOK).JSON(response)
}
