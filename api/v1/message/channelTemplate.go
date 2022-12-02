package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/common/response"
	"aixinge/api/model/message"
	"aixinge/api/model/validation"
	"aixinge/global"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ChannelTemplate struct {
}

// Create
// @Tags ChannelTemplate
// @Summary 创建渠道模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body message.ChannelTemplate true "创建渠道模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"渠道模板创建成功"}"
// @Router /v1/channel-template/create [post]
func (b *ChannelTemplate) Create(c *fiber.Ctx) error {
	var ct message.ChannelTemplate
	_ = c.BodyParser(&ct)
	err := channelTemplateService.Create(ct)
	if err != nil {
		global.LOG.Error("创建渠道模板失败！", zap.Any("err", err))
		return response.FailWithMessage("渠道模板创建失败："+err.Error(), c)
	}
	return response.OkWithMessage("渠道模板创建成功！", c)
}

// Delete
// @Tags ChannelTemplate
// @Summary 删除渠道模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/channel-template/delete [post]
func (b *ChannelTemplate) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := channelTemplateService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}
