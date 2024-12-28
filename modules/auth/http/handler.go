package http

import (
	"net/http"
	"sensor_monitoring_be/helper"
	"sensor_monitoring_be/modules/auth/service"
	"strconv"
	"time"

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
	valid, err := h.service.CheckPassword(c.Context(), credentials["username"], credentials["password"])
	if !valid || err != nil {
		return c.Status(http.StatusOK).JSON(helper.APIResponse("Login Failed, Invalid Credentials", http.StatusBadRequest, "Error", nil))
	}

	// Call AuthenticateService here and return the response
	resp, err := h.service.Authenticate(c.Context(), credentials["username"], credentials["password"])
	mappedResp := map[string]interface{}{
		"access_token":  resp[0]["access_token"],
		"clientId":      resp[0]["client_id"],
		"clientSecret":  resp[0]["client_secret"],
		"expires_in":    int(resp[0]["expires_at"].(time.Time).Unix()),
		"userId":        resp[0]["user_id"],
		"refresh_token": resp[0]["access_token"],
		"scope":         "tlinkAppId",
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal Server Error", http.StatusInternalServerError, "Error", nil))
	}
	return c.Status(http.StatusOK).JSON(helper.APIResponse("Login Success", http.StatusOK, "OK", mappedResp))
}

func (h *AuthHandler) HandleCreateUser(c *fiber.Ctx) error {
	var user map[string]interface{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}
	// Check first if username exist
	exist, err := h.service.CheckUsernameExist(c.Context(), user["username"].(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal Server Error", http.StatusInternalServerError, "Error", nil))
	}
	if exist {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Username already exist", http.StatusBadRequest, "Bad Request", nil))
	}
	// Call CreateUser service here and return the response
	resp, err := h.service.CreateUser(c.Context(), user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal Server Error", http.StatusInternalServerError, "Error", nil))
	}
	return c.Status(http.StatusOK).JSON(helper.APIResponse("User Created Success", http.StatusOK, "OK", resp))
}

func (h *AuthHandler) HandleEditUser(c *fiber.Ctx) error {
	var user map[string]interface{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid request", http.StatusBadRequest, "Bad Request", nil))
	}

	// Extract user ID from the request and ensure it's an integer
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid user ID", http.StatusBadRequest, "Bad Request", nil))
	}

	// Call EditUser service here and return the response
	resp, err := h.service.EditUser(c.Context(), userID, user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal Server Error", http.StatusInternalServerError, "Error", nil))
	}
	return c.Status(http.StatusOK).JSON(helper.APIResponse("User Updated Success", http.StatusOK, "OK", resp))
}

func (h *AuthHandler) HandleDeleteUser(c *fiber.Ctx) error {
	// Extract user ID from the request and ensure it's an integer
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid user ID", http.StatusBadRequest, "Bad Request", nil))
	}

	// Call DeleteUser service here and return the response
	err = h.service.DeleteUser(c.Context(), userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal Server Error", http.StatusInternalServerError, "Error", nil))
	}
	return c.Status(http.StatusOK).JSON(helper.APIResponse("User Deleted Success", http.StatusOK, "OK", nil))
}
