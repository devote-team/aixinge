package message

import (
	v1 "aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type MailLogRouter struct {
}

func (m *MailLogRouter) InitMailLogRouter(router fiber.Router) {
	mlRouter := router.Group("mail-log")
	var mlApi = v1.AppApi.MessageApi.MailLog
	{
		mlRouter.Post("delete", mlApi.Delete) // 删除
		mlRouter.Post("page", mlApi.Page)     // 分页
	}
}
