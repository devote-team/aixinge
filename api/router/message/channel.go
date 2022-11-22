package message

import (
	v1 "aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type ChannelRouter struct {
}

func (c *ChannelRouter) InitChannelRouter(router fiber.Router) (R fiber.Router) {
	channelRouter := router.Group("channel")
	var channelApi = v1.AppApi.MessageApi.Channel
	{
		channelRouter.Post("create", channelApi.Create) // 创建
		channelRouter.Post("get", channelApi.Get)       // 根据id获取消息渠道
	}
	return channelRouter
}
