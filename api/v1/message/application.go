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

// Delete
// @Tags Menu
// @Summary 删除应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/application/delete [post]
func (b *Application) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := applicationService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags Menu
// @Summary 更新应用信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body message.Application true "应用信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/application/update [post]
func (b *Application) Update(c *fiber.Ctx) error {
	var app message.Application
	_ = c.BodyParser(&app)
	if err := validation.Verify(app, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	err, app := applicationService.Update(app)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	}

	return response.OkWithDetailed(messageRes.AppResponse{Application: app}, "更新成功", c)
}

// Get
// @Tags Menu
// @Summary 根据id获取应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "应用ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/application/get [post]
func (b *Application) Get(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := validation.Verify(idInfo, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, app := applicationService.GetById(idInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(messageRes.AppResponse{Application: app}, "获取成功", c)
	}
}
