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

type MailTemplate struct {
}

// Create
// @Tags MailTemplate
// @Summary 创建邮件模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body message.MailTemplate true "创建邮件模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"邮件模板创建成功"}"
// @Router /v1/mail-template/create [post]
func (b *MailTemplate) Create(c *fiber.Ctx) error {
	var mt message.MailTemplate
	_ = c.BodyParser(&mt)
	err := mailTemplateService.Create(mt)
	if err != nil {
		global.LOG.Error("创建邮件模板失败！", zap.Any("err", err))
		return response.FailWithMessage("邮件模板创建失败："+err.Error(), c)
	}
	return response.OkWithMessage("邮件模板创建成功！", c)
}

// Get
// @Tags MailTemplate
// @Summary 根据id获取邮件模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "邮件模板ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/mail-template/get [post]
func (b *MailTemplate) Get(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, mt := mailTemplateService.GetById(idInfo.ID); err != nil {
		global.LOG.Error("获取邮件模板失败!", zap.Any("err", err))
		return response.FailWithMessage("获取邮件模板失败", c)
	} else {
		return response.OkWithDetailed(messageRes.MailTemplateResponse{MailTemplate: mt}, "获取邮件模板成功", c)
	}
}
