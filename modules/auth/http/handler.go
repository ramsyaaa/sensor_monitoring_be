package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sensor_monitoring_be/helper"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) HandleAuth(c *fiber.Ctx) error {
	baseURL := os.Getenv("BASE_URL")
	authHeader := os.Getenv("AUTH_HEADER")
	var credentials map[string]string
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	username := credentials["username"]
	password := credentials["password"]

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token?grant_type=password&username=%s&password=%s", baseURL, username, password), nil)
	req.Header.Set("Authorization", authHeader)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var responseMap map[string]interface{}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, &responseMap)
	if err != nil {
		return err
	}

	response := helper.APIResponse("Login Success", http.StatusOK, "OK", responseMap)
	return c.Status(http.StatusOK).JSON(response)
}
