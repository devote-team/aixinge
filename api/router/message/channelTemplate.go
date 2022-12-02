package message

import (
	v1 "aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type ChannelTemplateRouter struct {
}

func (c *ChannelRouter) InitChannelTemplateRouter(router fiber.Router) {
	ctRouter := router.Group("channel-template")
	var ctApi = v1.AppApi.MessageApi.ChannelTemplate
	{
		ctRouter.Post("create", ctApi.Create) // 创建
		ctRouter.Post("delete", ctApi.Delete) // 删除
	}
}
