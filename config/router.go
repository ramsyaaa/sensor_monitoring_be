package config

import (
	"log"
	"os"
	"sensor_monitoring_be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Route() {

	app := fiber.New()
	// Use the cors middleware to allow all origins and methods
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Create a new Fiber app for the "api/v1" prefix group
	api := fiber.New()

	routes.AuthRouter(api)
	routes.DeviceRouter(api)

	// Mount the "api/v1" group under the main app
	app.Mount("/api/v1", api)

	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}
