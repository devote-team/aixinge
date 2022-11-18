package message

import (
	"aixinge/api/model/common/response"
	"aixinge/api/model/message"
	"aixinge/global"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Application struct {
}

// Create
// @Tags Application
// @Summary 创建应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body message.Application true "创建应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"应用创建成功"}"
// @Router /v1/application/create [post]
func (b *Application) Create(c *fiber.Ctx) error {
	var app message.Application
	_ = c.BodyParser(&app)
	err := applicationService.Create(app)
	if err != nil {
		global.LOG.Error("创建应用失败！", zap.Any("err", err))
		return response.FailWithMessage("应用创建失败："+err.Error(), c)
	}
	return response.OkWithMessage("应用创建成功！", c)
}
