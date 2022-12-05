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
		appRouter.Post("delete", appApi.Delete) // 删除应用
		appRouter.Post("update", appApi.Update) // 更新应用信息
		appRouter.Post("get", appApi.Update)    // 根据id获取应用
	}
}
