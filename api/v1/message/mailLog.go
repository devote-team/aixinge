package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/common/response"
	"aixinge/api/model/validation"
	"aixinge/global"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type MailLog struct {
}

// Delete
// @Tags MailLog
// @Summary 删除邮件日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID集合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/mail-log/delete [post]
func (m *MailLog) Delete(c *fiber.Ctx) error {
	var idsReq request.IdsReq
	_ = c.BodyParser(&idsReq)
	if err := validation.Verify(idsReq, validation.Id); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := mailLogService.Delete(idsReq); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// Page
// @Tags MailLog
// @Summary 分页获取邮件日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/mail-log/page [post]
func (m *MailLog) Page(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := validation.Verify(pageInfo, validation.PageInfo); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := mailLogService.Page(pageInfo); err != nil {
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
