package http

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, handler *AuthHandler) {
	app.Post("/authenticate", handler.HandleAuth)
	app.Post("/user/create", handler.HandleCreateUser)
	app.Put("/user/edit/:id", handler.HandleEditUser)
	app.Delete("/user/delete/:id", handler.HandleDeleteUser)
	app.Get("/user/list", handler.HandleListUsers)
	app.Get("/refresh-token", handler.HandleRefreshToken)
}
