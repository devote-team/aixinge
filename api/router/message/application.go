package message

import (
	v1 "aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type ApplicationRouter struct {
}

func (a *ApplicationRouter) InitApplicationRouter(router fiber.Router) {
	appRouter := router.Group("app")
	var appApi = v1.AppApi.MessageApi.Application
	{
		appRouter.Post("create", appApi.Create) // 创建
	}
}
