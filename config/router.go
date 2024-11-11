package config

import (
	"log"
	"os"
	"sensor_monitoring_be/helper" // Import your custom helper package
	"sensor_monitoring_be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) {

	// Create a new Fiber app instance
	app := fiber.New()

	// Use the cors middleware to allow all origins and methods
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	app.Static("/static", "./static")

	// Serve the HTML dashboard on the root path
	app.Get("/log-viewer", func(c *fiber.Ctx) error {
		return c.SendFile("./static/index.html")
	})

	// Get available log files (uses helper function)
	app.Get("/logs", helper.GetLogFiles)

	// Get content from a specific log file (uses helper function)
	app.Get("/logs/:filename", helper.GetLogFileContent)

	// Use the custom logger middleware only for API routes

	// Create a new Fiber app for the "api/v1" prefix group
	api := fiber.New()
	api.Use(helper.LogToFile())

	// Set up your routes
	routes.AuthRouter(api)
	routes.DeviceRouter(api)
	routes.GeoMappingRouter(api, db)

	// Mount the "api/v1" group under the main app
	app.Mount("/api/v1", api)

	// Start the server on the specified port (from the environment variable)
	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}
