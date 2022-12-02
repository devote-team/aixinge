package system

import (
	"aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(router fiber.Router) fiber.Router {
	var userApi = v1.AppApi.SystemApi.User
	{
		router.Post("login", userApi.Login)
		router.Post("refresh-token", userApi.RefreshToken)
	}
	return router
}
