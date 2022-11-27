package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() fiber.Handler {
	config := logger.ConfigDefault
	config.Format = "${time} ${status} - ${latency} ${method} ${path} \n"
	config.TimeFormat = "2006/01/02 - 15:04:05"
	return logger.New(config)

}
