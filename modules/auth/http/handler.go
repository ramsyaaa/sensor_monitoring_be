package http

import (
	"net/http"
	"sensor_monitoring_be/helper"
	"sensor_monitoring_be/modules/auth/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}
func (h *AuthHandler) HandleAuth(c *fiber.Ctx) error {
	var credentials map[string]string
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	if credentials["username"] != "bpbd" || credentials["password"] != "b9d6devel" {
		return c.Status(http.StatusOK).JSON(helper.APIResponse("Login Failed, Invalid Credentials", http.StatusBadRequest, "Error", nil))
	}

	// Call AuthenticateService here and return the response
	resp, err := h.service.Authenticate(c.Context(), credentials["username"], credentials["password"])
	mappedResp := map[string]interface{}{
		"access_token":  resp[0]["access_token"],
		"clientId":      resp[0]["client_id"],
		"clientSecret":  resp[0]["client_secret"],
		"expires_in":    resp[0]["expires_at"],
		"userId":        resp[0]["user_id"],
		"refresh_token": resp[0]["access_token"],
		"scope":         "tlinkAppId",
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal Server Error", http.StatusInternalServerError, "Error", nil))
	}
	return c.Status(http.StatusOK).JSON(helper.APIResponse("Login Success", http.StatusOK, "OK", mappedResp))
}
