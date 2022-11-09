package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Cors 处理跨域请求,支持options访问
func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods:     "POST,GET,OPTIONS,DELETE,PUT",
		AllowHeaders:     "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type",
		MaxAge:           0,
	})
}
