package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/common/response"
	"aixinge/api/model/message"
	messageRes "aixinge/api/model/message/response"
	"aixinge/api/model/validation"
	"aixinge/global"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Channel struct {
}

// Create
// @Tags Channel
// @Summary 创建消息渠道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body message.MailTemplate true "创建消息渠道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"消息渠道创建成功"}"
// @Router /v1/channel/create [post]
func (b *Channel) Create(c *fiber.Ctx) error {
	var mt message.Channel
	_ = c.BodyParser(&mt)
	err := channelService.Create(mt)
	if err != nil {
		global.LOG.Error("创建消息渠道失败！", zap.Any("err", err))
		return response.FailWithMessage("消息渠道创建失败："+err.Error(), c)
	}
	return response.OkWithMessage("消息渠道创建成功！", c)
}

// Get
// @Tags Channel
// @Summary 根据id获取消息渠道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "消息渠道ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/channel/get [post]
func (b *Channel) Get(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, chl := channelService.GetById(idInfo.ID); err != nil {
		global.LOG.Error("获取消息渠道失败!", zap.Any("err", err))
		return response.FailWithMessage("获取消息渠道失败", c)
	} else {
		return response.OkWithDetailed(messageRes.ChannelResponse{Channel: chl}, "获取消息渠道成功", c)
	}
}
