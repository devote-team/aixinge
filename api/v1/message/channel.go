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
// @Param data body message.Channel true "创建消息渠道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"消息渠道创建成功"}"
// @Router /v1/channel/create [post]
func (b *Channel) Create(c *fiber.Ctx) error {
	var channel message.Channel
	_ = c.BodyParser(&channel)
	err := channelService.Create(channel)
	if err != nil {
		global.LOG.Error("创建消息渠道失败！", zap.Any("err", err))
		return response.FailWithMessage("消息渠道创建失败："+err.Error(), c)
	}
	return response.OkWithMessage("消息渠道创建成功！", c)
}

// Delete
// @Tags Channel
// @Summary 删除消息渠道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/channel/delete [post]
func (b *Channel) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := channelService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags Channel
// @Summary 更新消息渠道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body message.Channel true "消息渠道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /v1/channel/update [post]
func (b *Channel) Update(c *fiber.Ctx) error {
	var channel message.Channel
	_ = c.BodyParser(&channel)
	if err := validation.Verify(channel, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	err, channel := channelService.Update(channel)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	}

	return response.OkWithDetailed(messageRes.ChannelResponse{Channel: channel}, "更新成功", c)
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

// Page
// @Tags Channel
// @Summary 分页获取消息渠道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/channel/page [post]
func (b *Channel) Page(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := validation.Verify(pageInfo, validation.PageInfo); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := channelService.Page(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
