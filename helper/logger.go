package helper

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type LoggerConfig struct {
	Format       string
	TimeFormat   string
	TimeZone     string
	Output       *os.File
	CustomTags   map[string]logger.LogFunc
	Done         func(*fiber.Ctx, []byte)
	DisableColor bool
}

func NewLogger(config LoggerConfig) fiber.Handler {
	return logger.New(logger.Config{
		Format:     config.Format,
		TimeFormat: config.TimeFormat,
		TimeZone:   config.TimeZone,
		Output:     config.Output,
		CustomTags: config.CustomTags,
		Done:       config.Done,
	})
}

func DefaultLogger() fiber.Handler {
	return NewLogger(LoggerConfig{
		Format: "[${time}] ${ip} ${status} - ${method} ${path}\n",
	})
}

func RequestIDLogger() fiber.Handler {
	return requestid.New()
}

func CustomLogger(config LoggerConfig) fiber.Handler {
	return NewLogger(config)
}

func CustomFileLogger(filePath string) (*os.File, fiber.Handler, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file, NewLogger(LoggerConfig{
		Output: file,
	}), err
}

func CustomTagLogger(customTags map[string]logger.LogFunc) fiber.Handler {
	return NewLogger(LoggerConfig{
		CustomTags: customTags,
	})
}

func CallbackLogger(done func(*fiber.Ctx, []byte)) fiber.Handler {
	return NewLogger(LoggerConfig{
		Done: done,
	})
}

func DisableColorLogger() fiber.Handler {
	return NewLogger(LoggerConfig{
		DisableColor: true,
	})
}

func LogToFile() fiber.Handler {
	// Define log file path based on the current date
	filePath := "logs/" + time.Now().Format("20060102") + "-log.log"

	// Open the file with read/write, create or append mode
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Create a writer to write to the file only
	// Tidak ada lagi stdout karena perubahan untuk tidak mencetak log di command line

	// Return the fiber logger middleware, directing it to the file writer
	return logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Local",
		Output:     file, // Tidak ada lagi stdout karena perubahan untuk tidak mencetak log di command line
	})
}
