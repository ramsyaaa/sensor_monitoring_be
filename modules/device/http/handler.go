package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sensor_monitoring_be/helper"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type DeviceHandler struct {
}

func NewDeviceHandler() *DeviceHandler {
	return &DeviceHandler{}
}

func (h *DeviceHandler) HandleGetDevices(c *fiber.Ctx) error {
	tlinkAppId := os.Getenv("TLINK_APP_ID")
	if tlinkAppId == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	authHeader := c.Request().Header.Peek("Authorization")
	if string(authHeader) == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	var request map[string]int
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	userId := request["userId"]
	currPage := request["currPage"]
	pageSize := request["pageSize"]

	// Hit external API to get devices
	baseURL := os.Getenv("BASE_URL")
	reqBody := fmt.Sprintf(`{"userId":%d,"currPage":%d,"pageSize":%d}`, userId, currPage, pageSize)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/device/getDevices", baseURL), strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("tlinkAppId", tlinkAppId)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", string(authHeader))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 {
		return c.Status(http.StatusOK).JSON(helper.APIResponse("No devices found", http.StatusOK, "OK", nil))
	}
	var devicesResponse struct {
		CurrPage int                      `json:"currPage"`
		Pages    int                      `json:"pages"`
		DataList []map[string]interface{} `json:"dataList"`
		PageSize int                      `json:"pageSize"`
		RowCount int                      `json:"rowCount"`
	}
	err = json.Unmarshal(respBody, &devicesResponse)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to unmarshal devices", http.StatusBadRequest, "Bad Request", nil))
	}

	response := helper.APIResponse("Devices fetched successfully", http.StatusOK, "OK", devicesResponse)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *DeviceHandler) HandleGetDevicesWithSensor(c *fiber.Ctx) error {
	tlinkAppId := os.Getenv("TLINK_APP_ID")
	if tlinkAppId == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	authHeader := c.Request().Header.Peek("Authorization")
	if string(authHeader) == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	var request map[string]int
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	userId := request["userId"]
	currPage := request["currPage"]
	pageSize := request["pageSize"]

	// Hit external API to get devices
	baseURL := os.Getenv("BASE_URL")
	reqBody := fmt.Sprintf(`{"userId":%d,"currPage":%d,"pageSize":%d}`, userId, currPage, pageSize)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/device/getDeviceSensorDatas", baseURL), strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("tlinkAppId", tlinkAppId)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", string(authHeader))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 {
		return c.Status(http.StatusOK).JSON(helper.APIResponse("No devices found", http.StatusOK, "OK", nil))
	}
	var devicesResponse struct {
		CurrPage int                      `json:"currPage"`
		Pages    int                      `json:"pages"`
		DataList []map[string]interface{} `json:"dataList"`
		PageSize int                      `json:"pageSize"`
		RowCount int                      `json:"rowCount"`
	}
	err = json.Unmarshal(respBody, &devicesResponse)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to unmarshal devices", http.StatusBadRequest, "Bad Request", nil))
	}

	response := helper.APIResponse("Devices With Sensor Data fetched successfully", http.StatusOK, "OK", devicesResponse)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *DeviceHandler) HandleGetSingleDevice(c *fiber.Ctx) error {
	tlinkAppId := os.Getenv("TLINK_APP_ID")
	if tlinkAppId == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	authHeader := c.Request().Header.Peek("Authorization")
	if string(authHeader) == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	var request map[string]int
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	userId := request["userId"]
	deviceId := request["deviceId"]
	currPage := request["currPage"]
	pageSize := request["pageSize"]

	// Hit external API to get devices
	baseURL := os.Getenv("BASE_URL")
	reqBody := fmt.Sprintf(`{"userId":%d,"deviceId":%d,"currPage":%d,"pageSize":%d}`, userId, deviceId, currPage, pageSize)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/device/getSingleDeviceDatas", baseURL), strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("tlinkAppId", tlinkAppId)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", string(authHeader))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 {
		return c.Status(http.StatusOK).JSON(helper.APIResponse("No devices found", http.StatusOK, "OK", nil))
	}
	var devicesResponse struct {
		CurrPage int         `json:"currPage"`
		Pages    int         `json:"pages"`
		PageSize int         `json:"pageSize"`
		RowCount int         `json:"rowCount"`
		Device   interface{} `json:"device"`
	}
	err = json.Unmarshal(respBody, &devicesResponse)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to unmarshal devices", http.StatusBadRequest, "Bad Request", nil))
	}

	response := helper.APIResponse("Devices Single Data fetched successfully", http.StatusOK, "OK", devicesResponse)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *DeviceHandler) HandleGetSingleSensor(c *fiber.Ctx) error {
	tlinkAppId := os.Getenv("TLINK_APP_ID")
	if tlinkAppId == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	authHeader := c.Request().Header.Peek("Authorization")
	if string(authHeader) == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	var request map[string]int
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	userId := request["userId"]
	sensorId := request["sensorId"]

	// Hit external API to get devices
	baseURL := os.Getenv("BASE_URL")
	reqBody := fmt.Sprintf(`{"userId":%d,"sensorId":%d}`, userId, sensorId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/device/getSingleSensorDatas", baseURL), strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("tlinkAppId", tlinkAppId)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", string(authHeader))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 {
		return c.Status(http.StatusOK).JSON(helper.APIResponse("No devices found", http.StatusOK, "OK", nil))
	}
	var responseMap map[string]interface{}
	err = json.Unmarshal(respBody, &responseMap)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to unmarshal devices", http.StatusBadRequest, "Bad Request", nil))
	}

	response := helper.APIResponse("Devices Single Data fetched successfully", http.StatusOK, "OK", responseMap)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *DeviceHandler) HandleGetSensorHistory(c *fiber.Ctx) error {
	tlinkAppId := os.Getenv("TLINK_APP_ID")
	if tlinkAppId == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	authHeader := c.Request().Header.Peek("Authorization")
	if string(authHeader) == "" {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Unauthorized", nil))
	}

	type RequestBody struct {
		UserId      int    `json:"userId"`
		SensorId    int    `json:"sensorId"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
		PageSize    int    `json:"pageSize"`
		PagingState string `json:"pagingState"`
	}

	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	userId := requestBody.UserId
	sensorId := requestBody.SensorId
	startDate := requestBody.StartDate
	endDate := requestBody.EndDate
	pageSize := requestBody.PageSize
	pagingState := requestBody.PagingState
	// Hit external API to get devices
	baseURL := os.Getenv("BASE_URL")
	reqBody := fmt.Sprintf(`{"userId":%d,"sensorId":%d,"startDate":"%s","endDate":"%s","pageSize":%d,"pagingState":"%s"}`, userId, sensorId, startDate, endDate, pageSize, pagingState)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/device/getSensorHistroy", baseURL), strings.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("tlinkAppId", tlinkAppId)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", string(authHeader))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 {
		return c.Status(http.StatusOK).JSON(helper.APIResponse("No devices found", http.StatusOK, "OK", nil))
	}
	var responseMap map[string]interface{}
	err = json.Unmarshal(respBody, &responseMap)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to unmarshal devices", http.StatusBadRequest, "Bad Request", nil))
	}

	response := helper.APIResponse("Sensor History data fetched successfully", http.StatusOK, "OK", responseMap)
	return c.Status(http.StatusOK).JSON(response)
}
