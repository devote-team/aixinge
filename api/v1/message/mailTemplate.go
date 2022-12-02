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
func (e *MailTemplate) Create(c *fiber.Ctx) error {
	var mt message.MailTemplate
	_ = c.BodyParser(&mt)
	err := mailTemplateService.Create(mt)
	if err != nil {
		global.LOG.Error("创建邮件模板失败！", zap.Any("err", err))
		return response.FailWithMessage("邮件模板创建失败："+err.Error(), c)
	}
	return response.OkWithMessage("邮件模板创建成功！", c)
}

// Delete
// @Tags MailTemplate
// @Summary 删除邮件模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/mail-template/delete [post]
func (e *MailTemplate) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := mailTemplateService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags MailTemplate
// @Summary 更新邮件模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body message.MailTemplate true "邮件模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /v1/mail-template/update [post]
func (e *MailTemplate) Update(c *fiber.Ctx) error {
	var mt message.MailTemplate
	_ = c.BodyParser(&mt)
	if err := validation.Verify(mt, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	err, mt := mailTemplateService.Update(mt)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	}

	return response.OkWithDetailed(messageRes.MailTemplateResponse{MailTemplate: mt}, "更新成功", c)
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
func (e *MailTemplate) Get(c *fiber.Ctx) error {
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

// Page
// @Tags MailTemplate
// @Summary 分页获取邮件模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/mail-template/page [post]
func (e *MailTemplate) Page(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := validation.Verify(pageInfo, validation.PageInfo); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := mailTemplateService.Page(pageInfo); err != nil {
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
